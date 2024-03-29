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
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	address  string
	services sync.Map
	mu       sync.Mutex
	m        *middle
	lis      net.Listener
}

var (
	DefaultServer  = NewServer(DefaultAddress)
	DefaultAddress = ":9115"
)

var (
	procSignalChan = make(chan os.Signal)
	GOEXIT         = make(chan struct{}, 1)
)

func NewServer(address string) *Server {
	if address == "" {
		address = fmt.Sprintf("%s:%d", config.GetGrpcServerConfig().Addr, config.GetGrpcServerConfig().Port)
	}
	return &Server{address: address, m: newMiddle()}
}

func Run() {
	DefaultServer.Run()
}

func Close() error {
	return DefaultServer.Close()
}

func (s *Server) Close() error {
	return s.lis.Close()
}

func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}

func RegisterMiddleware(mf ...MiddleFunc) {
	DefaultServer.RegisterMiddleware(mf...)
}

func (s *Server) RegisterMiddleware(mf ...MiddleFunc) {
	s.m.f = append(s.m.f, mf...)
}

func (s *Server) Register(rcvr interface{}) error {
	src := newService(rcvr)
	if _, dup := s.services.LoadOrStore(src.name, src); dup {
		return errors.New("rpc: service already defined: " + src.name)
	}
	log.Println("RegisterMiddleware ok...............................")
	return nil
}

func (s *Server) Run() {

	listen, err := net.Listen("tcp4", s.address)
	if err != nil {
		panic(err)
	}
	s.lis = listen

	log.Printf("run at in address: %s.........................\n", s.address)
	go handleProcessSignal()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		Gresources().add(tools.UUID(), conn)
		go sendResponse("Welcome to the world of Red police", nil, conn)
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {

	defer func() {
		conn.Close()
		log.Println("over")
		if err := recover(); err != nil {
		}
	}()

	decoder := json.NewDecoder(conn)

L:
	for {
		var req Request
		begin := time.Now()
		s.setReadDeadline(begin, conn)
		if err := decoder.Decode(&req); err != nil {
			if err == io.EOF {
				break L
			}
			sendResponse(nil, err, conn)
			break L
		} else {
			go s.handleRequest(&req, conn)
		}
	}
}

func (s *Server) setReadDeadline(startTime time.Time, conn net.Conn) {
	conn.SetReadDeadline(startTime.Add(time.Second * time.Duration(config.GetGrpcServerConfig().ReadDeadLine)))
}

func (s *Server) handleRequest(req *Request, writer io.WriteCloser) {

	svc, mtype, serr := s.findService(req.ServiceMethod)
	if serr != nil {
		sendResponse(req, serr, writer)
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
		sendResponse(req, err, writer)
		return
	}
	req.MethodReflectData = farg

	if err := s.m.call(req); err != nil {
		sendResponse(nil, err, writer)
		return
	}
	if err := iReq.svc.call(iReq.mtype, reflect.ValueOf(farg), iReq.replyv); err != nil {
		sendResponse(nil, err, writer)
		return
	}

	sendResponse(iReq.replyv.Elem().Interface(), nil, writer)

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

func handleProcessSignal() {
	var sig os.Signal
	signal.Notify(
		procSignalChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)
	for {
		sig = <-procSignalChan
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGUSR1:
			GOEXIT <- struct{}{}
			return
		default:
		}
	}
}
