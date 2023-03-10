package network

import (
	"20red_police/config"
	"20red_police/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Server struct {
	address  string
	services sync.Map
	conns    map[string]*net.Conn
	mu       sync.Mutex
}

var (
	DefaultServer  = NewServer(DefaultAddress)
	DefaultAddress = ":9114"
)

func NewServer(address string) *Server {
	if address == "" {
		address = fmt.Sprintf("%s:%d", config.GetGrpcServerConfig().Addr, config.GetGrpcServerConfig().Port)
	}
	return &Server{address: address, conns: map[string]*net.Conn{}}
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

	listen, err := net.Listen("tcp4", s.address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("run at in address: %s.........................\n", s.address)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	threadID := tools.UUID()
	s.mu.Lock()
	s.conns[threadID] = &conn
	s.mu.Unlock()
	decoder := json.NewDecoder(conn)
L:
	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {
			s.mu.Lock()
			delete(s.conns, threadID)
			s.mu.Unlock()
			if err == io.EOF {
				break L
			}
			s.sendResponse(nil, err, conn)
			break L
		} else {
			go s.handleRequest(&req, conn)
		}
	}
}

func (s *Server) handleRequest(req *Request, writer io.WriteCloser) {

	//log.Printf("**************************[read data=%+v]**************************\n", *req)

	svc, mtype, serr := s.findService(req.ServiceMethod)
	if serr != nil {
		s.sendResponse(req, serr, writer)
		return
	}

	var iReq innerRequest
	iReq.svc = svc
	iReq.mtype = mtype
	iReq.argv = iReq.mtype.newArgv()
	iReq.replyv = iReq.mtype.newReplyv()
	argvi := iReq.argv.Interface()
	if iReq.argv.Type().Kind() != reflect.Ptr {
		argvi = iReq.argv.Addr().Interface() //函数返回一个持有指向v持有者的指针的Value封装。如果v.CanAddr()返回假，调用本方法会panic。Addr一般用于获取结构体字段的指针或者切片的元素（的Value封装）以便调用需要指针类型接收者的方法。
	}
	argvi = req.MetaData
	farg := iReq.mtype.newArgv().Interface()
	if err := mapstructure.Decode(argvi, &farg); err != nil {
		s.sendResponse(req, err, writer)
		return
	}
	if err := iReq.svc.call(iReq.mtype, reflect.ValueOf(farg), iReq.replyv); err != nil {
		s.sendResponse(nil, err, writer)
		return
	}

	s.sendResponse(iReq.replyv.Elem().Interface(), nil, writer)
}

func (s *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {

	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}
	serviceName := serviceMethod[:dot]
	methodName := serviceMethod[dot+1:]

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

func (s *Server) sendResponse(data interface{}, err error, writer io.WriteCloser) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	res := &Response{Data: data, Err: errMsg}
	if err = json.NewEncoder(writer).Encode(res); err != nil {
		log.Println("send response err:", err)
	}
}

func (s *Server) ping() {
	timer := time.NewTicker(time.Second * 2)
	go func() {
		for {
			select {
			case <-timer.C:
				for threadID, conn := range s.conns {
					_, err2 := (*conn).Write([]byte("abc\n"))
					if err2 == io.EOF {
						delete(s.conns, threadID)
						return
					}
				}
			default:

			}
		}
	}()
}
