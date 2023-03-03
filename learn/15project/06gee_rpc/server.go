package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec/codec"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

/*
*
客户端与服务端的通信需要协商一些内容，例如 HTTP 报文，分为 header 和 body 2 部分，body 的格式和长度通过 header 中的 Content-Type 和 Content-Length 指定，
服务端通过解析 header 就能够知道如何从 body 中读取需要的信息。对于 RPC 协议来说，这部分协商是需要自主设计的。
为了提升性能，一般在报文的最开始会规划固定的字节，来协商相关的信息。比如第1个字节用来表示序列化方式，第2个字节表示压缩方式，第3-6字节表示 header 的长度，7-10 字节表示 body 的长度。
对于 GeeRPC 来说，目前需要协商的唯一一项内容是消息的编解码方式。我们将这部分信息，放到结构体 Option 中承载。目前，已经进入到服务端的实现阶段了。
*/
const MagicNumber = 0x3bef5c

const (
	connected        = "200 Connnected to Gee RPC"
	defaultRPCPath   = "_geerpc_"
	defaultDebugPath = "/default/geerpc"
)

var invalidRequest = struct{}{}

type Option struct {
	MagicNumber    int
	CodecType      codec.Type
	ConnectTimeout time.Duration
	HandleTimeout  time.Duration
}

var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second * 10,
}

func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "CONNECT" {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "405 must CONNECT\n")
		return
	}

	//获取unix套接字
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Print("rpc hijacking ", req.RemoteAddr, ": ", err.Error())
		return
	}
	_, _ = io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
	server.ServeConn(conn)
}

func (server *Server) HandleHTTP() {
	http.Handle(defaultRPCPath, server)
	http.Handle(defaultDebugPath, debugHTTP{server})
	log.Println("rpc server debug path:", defaultDebugPath)
}

func HandleHTTP() {
	DefaultServer.HandleHTTP()
}

/**
一般来说，涉及协议协商的这部分信息，需要设计固定的字节来传输的。
但是为了实现上更简单，GeeRPC 客户端固定采用 JSON 编码 Option，后续的 header 和 body 的编码方式由 Option 中的 CodeType 指定，
服务端首先使用 JSON 解码 Option，然后通过 Option 的 CodeType 解码剩余的内容。即报文将以这样的形式发送：
Option{MagicNumber: xxx, CodecType: xxx} | Header{ServiceMethod ...} | Body interface{} |
| <------      固定 JSON 编码      ------>  | <-------   编码方式由 CodeType 决定   ------->|

在一次连接中，Option 固定在报文的最开始，Header 和 Body 可以有多个，即报文可能是这样的。
1| Option | Header1 | Body1 | Header2 | Body2 | ...

*/

/**
首先定义了结构体 Server，没有任何的成员字段。
实现了 Accept 方式，net.Listener 作为参数，for 循环等待 socket 连接建立，并开启子协程处理，处理过程交给了 ServerConn 方法。
DefaultServer 是一个默认的 Server 实例，主要为了用户使用方便。
*/

type Server struct {
	serviceMap sync.Map
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr)
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup {
		return errors.New("rpc: service already defined: " + s.name)
	}
	return nil
}

func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}

/*
*
findService 的实现看似比较繁琐，但是逻辑还是非常清晰的。
因为 ServiceMethod 的构成是 “Service.Method”，
因此先将其分割成 2 部分，第一部分是 Service 的名称，第二部分即方法名。
现在 serviceMap 中找到对应的 service 实例，再从 service 实例的 method 中，找到对应的 methodType。
*/
func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {

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
	svci, ok := server.serviceMap.Load(serviceName)
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

// 阻塞，等待客户端连接
func (server *Server) Accept(lis net.Listener) {
	for {
		coon, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(coon)
	}
}

// 暴露的门面，方便上游调用
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

/*
*
ServeConn 的实现就和之前讨论的通信过程紧密相关了，
首先使用 json.NewDecoder 反序列化得到 Option 实例，检查 MagicNumber 和 CodeType 的值是否正确。
然后根据 CodeType 得到对应的消息编解码器，接下来的处理交给 serverCodec。
*/
func (server *Server) ServeConn(coon io.ReadWriteCloser) {
	defer func() {
		_ = coon.Close()
	}()

	//1. 首先使用 json.NewDecoder 反序列化得到 Option 实例，检查 MagicNumber 和 CodeType 的值是否正确。
	var opt Option
	if err := json.NewDecoder(coon).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}

	//2. 然后根据 CodeType 得到对应的消息编解码器，接下来的处理交给 serverCodec。
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}

	//3. 交给serverCodec处理
	server.serverCodec(f(coon))
}

/*
*
serveCodec 的过程非常简单。主要包含三个阶段

读取请求 readRequest
处理请求 handleRequest
回复请求 sendResponse
之前提到过，在一次连接中，允许接收多个请求，即多个 request header 和 request body，因此这里使用了 for 无限制地等待请求的到来，直到发生错误（例如连接被关闭，接收到的报文有问题等），这里需要注意的点有三个：

handleRequest 使用了协程并发执行请求。
处理请求是并发的，但是回复请求的报文必须是逐个发送的，并发容易导致多个回复报文交织在一起，客户端无法解析。在这里使用锁(sending)保证。
尽力而为，只有在 header 解析失败时，才终止循环。
*/
func (server *Server) serverCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		//1. 读取请求信息
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			//2. 无效的请求  回复请求
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		//3. 合法的请求，处理请求
		go server.handleRequest(cc, req, sending, wg, 0)
	}
	wg.Wait()
	_ = cc.Close()
}

// 一次请求
type request struct {
	h      *codec.Header //header头
	argv   reflect.Value //参数
	replyv reflect.Value //回复
	mtype  *methodType   //方法类型
	svc    *service      //具体服务
}

// 读取请求的header头
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

/*
*
readRequest 方法中最重要的部分，
即通过 newArgv() 和 newReplyv() 两个方法创建出两个入参实例，
然后通过 cc.ReadBody() 将请求报文反序列化为第一个入参 argv，
在这里同样需要注意 argv 可能是值类型，也可能是指针类型，所以处理方式有点差异。
*/
func (server *Server) readRequest(cc codec.Codec) (*request, error) {

	//1.读取header信息
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}

	//2.传入header，构造request
	req := &request{h: h}

	//3. 查找相应的服务
	req.svc, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err
	}

	//4. 构造请求的参数，构造回复的参数
	req.argv = req.mtype.newArgv()
	req.replyv = req.mtype.newReplyv()

	//通过反射，构造出接收client的参数的变量，方便后续将client中传来的参数，读取至❤新变量中
	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		//函数返回一个持有指向v持有者的指针的Value封装。如果v.CanAddr()返回假，调用本方法会panic。Addr一般用于获取结构体字段的指针或者切片的元素（的Value封装）以便调用需要指针类型接收者的方法。
		argvi = req.argv.Addr().Interface()
	}
	//5. 读取body信息至参数中
	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err:", err)
		return req, err
	}
	return req, nil
}

// 发送响应
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()

	// 输出字节流
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

/*
*
这里需要确保 sendResponse 仅调用一次，因此将整个过程拆分为 called 和 sent 两个阶段，在这段代码中只会发生如下两种情况：
called 信道接收到消息，代表处理没有超时，继续执行 sendResponse。
time.After() 先于 called 接收到消息，说明处理已经超时，called 和 sent 都将被阻塞。在 case <-time.After(timeout) 处调用 sendResponse。
*/
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()
	called := make(chan struct{})
	sent := make(chan struct{})
	go func() {

		//1. todo 传入参数，调用具体的服务
		err := req.svc.call(req.mtype, req.argv, req.replyv)

		//代表调用服务结束
		called <- struct{}{}

		if err != nil {
			req.h.Error = err.Error()
			//2-1发送无效的请求结果至客户端
			server.sendResponse(cc, req.h, invalidRequest, sending)
			//代表发送服务结果结束
			sent <- struct{}{}
			return
		}
		//2-2. 发送正常响应结果至客户端
		server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
		//代表发送服务结束
		sent <- struct{}{}
	}()

	//todo 3-1. 上游没有设置超时限制，直接从2个管道读取数据，这里可能会一直阻塞
	if timeout == 0 {
		<-called
		<-sent
		return
	}

	//todo 3-2 上游设置了超时限制，使用select，监听是否超时，如果超时，直接发送超时响应
	select {
	case <-time.After(timeout):
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, sending)
	case <-called:
		<-sent
	}
}
