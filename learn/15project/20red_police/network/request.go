package network

import (
	"reflect"
)

type Request struct {
	ServiceMethod string      `json:"service_method"`
	MetaData      interface{} `json:"meta_data"`
	Err           error
}

type InnerRequest struct {
	argv   reflect.Value //参数
	replyv reflect.Value //回复
	mtype  *methodType   //方法类型
	svc    *service      //具体服务
}

func handleRequest(req *Request) {

}
