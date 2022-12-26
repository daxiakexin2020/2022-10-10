package error

import (
	"errors"
	perror "github.com/pkg/errors"
)

type GoError struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Err  chan error
}

//错误类型
//底层服务错误，例如  http 代理   etcd
//10001-12000   http代理错误
//12001-14000   grpc代理错误
//20001-22000  etcd、console错误
//30001-40000  db错误  redis、mysql
//40001-50000 业务错误，例如service等错误

//业务错误 ，例如  验证等处理
//自定义错误
//系统错误     code 500
//自定义+err
//ok 0

//定义错误
//格式化错误
//包装错误

//todo 不建议每一个错误，都使用一个错误码，那是面向错误编程，可以使用同一个错误码，表示一个范围之内的错误

const (
	SUCCESS = 0

	SYSTEM_ERR = 500

	HTTP_PROXY_ERR = 10001
	GRPC_PROXY_ERR = 12001

	ETCD_ERR    = 20001
	CONSOLE_ERR = 20002

	REDIS_ERR = 30001
	MYSQL_ERR = 30002

	SERVER_ERR = 40001

	UNKNOWN_ERR = 1001
)

func NewGoErr() *GoError {
	ech := make(chan error, 10)
	return &GoError{
		Code: SUCCESS,
		Msg:  "ok",
		Err:  ech,
	}
}

func (gr *GoError) WithSystemError(msg string) error {
	return gr.setErr(SYSTEM_ERR, msg)
}

func (gr *GoError) WithHttpProxyError(msg string) error {
	return gr.setErr(HTTP_PROXY_ERR, msg)
}

func (gr *GoError) setErr(code int, msg string) error {
	gr.Code = code
	gr.Msg = msg
	e := errors.New(msg)
	gr.Err <- e
	return e
}

func (gr *GoError) setWrapErr(err error, code int, msg string) error {
	gr.Code = code
	gr.Msg = msg
	e := perror.WithMessage(err, msg)
	gr.Err <- e
	return e
}

func (gr *GoError) WrapWithSystemError(err error, msg string) error {
	return gr.setWrapErr(err, SYSTEM_ERR, msg)
}

func (gr *GoError) WrapWithHttpProxyError(err error, msg string) error {
	return gr.setWrapErr(err, HTTP_PROXY_ERR, msg)
}

func (gr *GoError) Unwrap() error {
	if len(gr.Err) == 0 {
		return nil
	}
	e := <-gr.Err
	ue := perror.Unwrap(e)
	if ue == nil {
		return e
	}
	return ue
}
