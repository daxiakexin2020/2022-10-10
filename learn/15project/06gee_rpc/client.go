package geerpc

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec/codec"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

/**
todo 整体大致逻辑流程
	 =>1.初始化一个conn   net.Conn  net包
	    =>2.将初始化的好的conn，传入自定义的的newClient中，初始化出来一个封装好的client
		=>3. 初始化一个call，封装具体要发送的class信息
		=>4. 使用client进行发送call

todo 获取client流程
	XDial=>DialHTTP->dialTimeout【这里通过传入回调函数实现，类似依赖注入，传入一个函数类型的函数，代为执行】
	->NewHTTPClient【如果是http类型的，直接给dialTimeout传入NewHTTPClient，依赖倒置】
	->NewClient->newClientCodec 【实例化出来client，并且开启协程，接收响应  receive()】

todo 发送call流程
*/

/**
Call 的设计
对 net/rpc 而言，一个函数需要能够被远程调用，需要满足如下五个条件：

the method’s type is exported. 方法类型可导出
the method is exported.     方法可导出
the method has two arguments, both exported (or builtin) types.  该方法有两个参数，都是导出(或内置)类型
the method’s second argument is a pointer.    第二参数是一个指针类型
the method has return type error. 该方法有返回类型错误。

更直观一些 func (t *T) MethodName(argType T1, replyType *T2) error
todo 根据上述要求，首先我们封装了结构体 Call 来承载一次 RPC 调用所需要的信息。

*/

/*
*
TODO 注意区分Call与Client的区别与联系，Call是代表一次请求，Client是发起请求的类，可以复用
*/
type Call struct {
	Seq           uint64      //请求编号，每个请求拥有唯一的编号
	ServiceMethod string      //服务方法名
	Args          interface{} //参数
	Reply         interface{} //返回值
	Error         error       //错误
	Done          chan *Call  //结束管道
}

/*
*
Client 的字段比较复杂：

cc 是消息的编解码器，和服务端类似，用来序列化将要发送出去的请求，以及反序列化接收到的响应。
sending 是一个互斥锁，和服务端类似，为了保证请求的有序发送，即防止出现多个请求报文混淆。
header 是每个请求的消息头，header 只有在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在 Client 结构体中可以复用。
seq 用于给发送的请求编号，每个请求拥有唯一编号。
pending 存储未处理完的请求，键是编号，值是 Call 实例。
closing 和 shutdown 任意一个值置为 true，则表示 Client 处于不可用的状态，但有些许的差别，closing 是用户主动关闭的，即调用 Close 方法，而 shutdown 置为 true 一般是有错误发生。
*/
type Client struct {
	cc       codec.Codec      //消息解码器，和服务端类似，用来序列化要发出去的请求，以及反序列化接收到的响应
	opt      *Option          //选项
	sending  sync.Mutex       //互斥锁，保证请求的有序发送，即防止出现多个请求报文混淆
	header   codec.Header     //每个请求的的消息头，heaeder只在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在Client中可以复用
	mu       sync.Mutex       //互斥锁
	seq      uint64           //请求编号，每个请求拥有唯一的编号
	pending  map[uint64]*Call //存储未处理的请求，键是编号，值是Call实例，Call代表的是一次请求
	closing  bool             //用户是否主动关闭  调用Close方法，会将此属性置为true
	shutdown bool             //有错误发生，会将此属性置为true
}

type clientResult struct {
	client *Client
	err    error
}

// conn客户端一般是接受连接   net.Conn  设计成一个函数类型，可以方便扩展，传入不同的类型，比如http，tcp的
type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

var _ io.Closer = (*Client)(nil)
var ErrShutdown = errors.New("connection is shut down")

func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.closing && !client.shutdown
}

// 创建 Client 实例时，首先需要完成一开始的协议交换，即发送 Option 信息给服务端。协商好消息的编解码方式之后，todo 再创建一个子协程调用 receive() 接收响应。
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

/**
对一个客户端端来说，接收响应、发送请求是最重要的 2 个功能。那么首先实现接收功能，接收到的响应有三种情况：
call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
call 存在，但服务端处理出错，即 h.Error 不为空。
call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。
*/

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

/**
call相关的3个方法
*/

// 将参数 call 添加到 client.pending 中，并更新 client.seq
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	call.Seq = client.seq //client复用的体现
	client.pending[call.Seq] = call
	client.seq++ //下一个call使用的seq
	return call.Seq, nil
}

// 根据 seq，从 client.pending 中移除对应的 call，并返回
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// 服务端或客户端发生错误时调用，将 shutdown 设置为 true，且将错误信息通知所有 pending 状态的 call。
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// 接收响应，反序列化结果
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		call := client.removeCall(h.Seq)
		switch {
		case call == nil:
			err = client.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	client.terminateCalls(err)
}

// 发送请求，接收一个call，发送
func (client *Client) send(call *Call) {
	client.sending.Lock()
	defer client.sending.Unlock()

	//1 将call先注册在pending中，待处理的队列中
	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	//2 设置请求需要的header头
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	//3 使用client发送请求
	if err = client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq) //从pending队列中，将此call移除
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// 创建一个client 带超时机制的client
func Dial(network string, address string, opts ...*Option) (client *Client, err error) {
	return dialTimeout(NewClient, network, address, opts...)
}

/*
在这里实现了一个超时处理的外壳 dialTimeout，这个壳将 NewClient 作为入参，在 2 个地方添加了超时处理的机制。
将 net.Dial 替换为 net.DialTimeout，如果连接创建超时，将返回错误。
2）使用子协程执行 NewClient，执行完成后则通过信道 ch 发送结果，如果 time.After() 信道先接收到消息，则说明 NewClient 执行超时，返回错误。
*/
func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {

	//1 解析发送的选项
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}

	//2 建立tcp/unix等连接
	conn, err := net.DialTimeout(network, address, opt.ConnectTimeout)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	ch := make(chan clientResult)

	//3 开启协程，处理 f是定义的类型函数
	go func() {
		client, err := f(conn, opt)
		ch <- clientResult{client: client, err: err}
	}()

	//4-1 没有设置超时限制，直接从ch中取数据，result := <-ch 可能会阻塞
	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}

	//4-2 设置了超时限制，使用select轮询超级ch与结果ch
	select {
	case <-time.After(opt.ConnectTimeout):
		return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
	case result := <-ch:
		return result.client, result.err
	}
}

/*
Go 和 Call 是客户端暴露给用户的两个 RPC 服务调用接口，Go 是一个异步接口，返回 call 实例。
Call 是对 Go 的封装，阻塞 call.Done，等待响应返回，是一个同步接口。
至此，一个支持异步和并发的 GeeRPC 客户端已经完成。
*/
func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

func (client *Client) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1))

	//加入超时机制，由上游ctx决定
	select {
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}

// 为了支持异步调用，Call 结构体中添加了一个字段 Done，Done 的类型是 chan *Call，当调用结束时，会调用 call.done() 通知调用方。
func (call *Call) done() {
	call.Done <- call
}

// http类型的conn，NewHTTPClient 是newClientFunc类型
func NewHTTPClient(conn net.Conn, opt *Option) (*Client, error) {
	_, _ = io.WriteString(conn, fmt.Sprintf("CONNECT %s HTTP/1.0\n\n", defaultRPCPath))
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
	if err == nil && resp.Status == connected {
		return NewClient(conn, opt)
	}
	if err == nil {
		err = errors.New("unexpected HTTP response " + resp.Status)
	}
	return nil, err
}

func DialHTTP(network string, address string, opts ...*Option) (*Client, error) {
	return dialTimeout(NewHTTPClient, network, address, opts...)
}

func XDial(rpcAddr string, opts ...*Option) (*Client, error) {
	parts := strings.Split(rpcAddr, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("rpc client err: wrong format '%s', expect protocol@addr", rpcAddr)
	}
	protocol := parts[0]
	addr := parts[1]
	switch protocol {
	case "http":
		return DialHTTP("tcp", addr, opts...)
	default:
		return Dial(protocol, addr, opts...)
	}
}
