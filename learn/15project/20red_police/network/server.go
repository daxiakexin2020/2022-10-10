package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"net"
	"reflect"
	"strings"
	"sync"
)

type Server struct {
	Address  string
	services sync.Map
}

var (
	DefaultServer  = NewServer(DefaultAddress)
	DefaultAddress = ":9114"
)

func NewServer(address string) *Server {
	return &Server{Address: address}
}

func Run() {
	DefaultServer.Run()
}

func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}

func (s *Server) Register(rcvr interface{}) error {
	src := newService(rcvr)
	if _, dup := s.services.LoadOrStore(src.name, src); dup {
		return errors.New("rpc: service already defined: " + src.name)
	}
	return nil
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	decoder := json.NewDecoder(conn)
	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {
			req.Err = err
			go s.handleRequest(&req, conn)
			continue
		} else {
			go s.handleRequest(&req, conn)
		}
	}
}

// {"service_method":"Server.Register","meta_data":"1"}
// {"service_method":"Server.Register","meta_data":{"name":"zz"}}
// {"service_method":"Server.Register2","meta_data":"kx"}

func (s *Server) handleRequest(req *Request, writer io.Writer) {
	fmt.Printf("read data=%+v\n", req)
	if req.Err != nil {
		s.sendResponse(req, req.Err, writer)
	} else {
		var iReq InnerRequest
		svc, mtype, err2 := s.findService(req.ServiceMethod)
		if err2 != nil {
			s.sendResponse(req, err2, writer)
		} else {
			iReq.svc = svc
			iReq.mtype = mtype
			iReq.argv = iReq.mtype.newArgv()
			iReq.replyv = iReq.mtype.newReplyv()
			//通过反射，构造出接收client的参数的变量，方便后续将client中传来的参数，读取至❤新变量中
			argvi := iReq.argv.Interface()
			if iReq.argv.Type().Kind() != reflect.Ptr {
				//函数返回一个持有指向v持有者的指针的Value封装。如果v.CanAddr()返回假，调用本方法会panic。Addr一般用于获取结构体字段的指针或者切片的元素（的Value封装）以便调用需要指针类型接收者的方法。
				argvi = iReq.argv.Addr().Interface()
			}
			argvi = req.MetaData
			farg := iReq.mtype.newArgv().Interface()
			if err := mapstructure.Decode(argvi, &farg); err != nil {
				s.sendResponse(nil, err, writer)
			}
			err := iReq.svc.call(iReq.mtype, reflect.ValueOf(farg), iReq.replyv)
			fmt.Println("err:", err, iReq.replyv.Elem().Interface())
			s.sendResponse(iReq.replyv.Elem().Interface(), nil, writer)
		}
	}
}

func (s *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {

	//1. 查找最后的.出现位置，切割出服务名、方法名
	//   serviceMethod  格式   Foo.bar   （服务名.方法名）
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}
	serviceName := serviceMethod[:dot]
	methodName := serviceMethod[dot+1:]

	//2. 通过服务名载入服务
	svci, ok := s.services.Load(serviceName)
	if !ok {
		err = errors.New("rpc server: can't find service " + serviceName)
		return
	}
	svc = svci.(*service)
	mtype = svc.method[methodName]
	if mtype == nil {
		err = errors.New("rpc server: can't find method " + methodName)
	}
	return
}

func (s *Server) sendResponse(data interface{}, err error, writer io.Writer) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	res := &Response{
		Data: data,
		Msg:  "ok",
		Err:  errMsg,
	}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		fmt.Println("send response err:", err)
	} else {
		fmt.Println("send response ok")
	}
}
