package network

import (
	"errors"
	"reflect"
)

type Request struct {
	ServiceMethod string      `json:"service_method"`
	MetaData      interface{} `json:"meta_data"`
	Header        *Header     `json:"header"`
}

type Header struct {
	Token string `json:"token"`
	BName string `json:"bname"`
}

func (r *Request) CheckHeader() error {
	if r.Header.Token == "" || r.Header.BName == "" {
		return errors.New("BName or Token is Needed")
	}
	return nil
}

type innerRequest struct {
	argv   reflect.Value //参数
	replyv reflect.Value //回复
	mtype  *methodType   //方法类型
	svc    *service      //具体服务
}
