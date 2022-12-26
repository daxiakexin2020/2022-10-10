package service

import (
	serror "14gateway/error"
	"14gateway/helper"
	proxy_http "14gateway/proxy/http"
	"14gateway/server_govern"
	"github.com/gin-gonic/gin"
	"net/http"
)

//解析请求
//发现注册服务
//调用远程服务

type clientServer struct {
	ctx     *gin.Context
	CMethod string
	CPath   string
	CAPi    string
	Addrs   []string
	ReqData interface{}
	GoErr   *serror.GoError
}

func NewClientServer(ctx *gin.Context) *clientServer {
	return &clientServer{ctx: ctx, GoErr: serror.NewGoErr()}
}

func (cs *clientServer) parseClientRequest() error {
	cmethod, cmexists := helper.GetCMethod(cs.ctx)
	if !cmexists || cmethod == "" {
		return cs.GoErr.WithHttpProxyError("不存在合法的请求方式")
	}

	cpath, cpexists := helper.GetCPath(cs.ctx)
	if !cpexists {
		return cs.GoErr.WithHttpProxyError("不存在合法的请求path" + cpath)
	}
	capi, caexists := helper.GetCApi(cs.ctx)
	if !caexists || capi == "" {
		return cs.GoErr.WithHttpProxyError("不存在合法的请求api")
	}
	crd, ccexists := helper.GetCReqData(cs.ctx)
	if ccexists && crd != nil {
		cs.ReqData = crd
	}
	cs.CMethod = cmethod
	cs.CPath = cpath
	cs.CAPi = capi
	return nil
}

func (cs *clientServer) discoveryServer() error {

	cmethod := server_govern.GetMethod(cs.CMethod)
	if cmethod == "" {
		return cs.GoErr.WithHttpProxyError("远程服务不支持此种请求方式 : " + cs.CMethod)
	}
	server := server_govern.NewServer(cs.CAPi, server_govern.WithMethodType(cmethod))
	addrs, err := server.Discovery()
	if err != nil {
		return err
	}
	if len(addrs) == 0 {
		return cs.GoErr.WithHttpProxyError("没有远程服务可以被使用")
	}

	cs.Addrs = addrs
	return nil
}

func (cs *clientServer) balance() (string, error) {
	index, err := helper.RandInt(len(cs.Addrs))
	if err != nil {
		return "", err
	}
	return cs.Addrs[index], nil
}

func (cs *clientServer) request() (*http.Response, error) {

	if len(cs.Addrs) == 0 {
		return nil, cs.GoErr.WithHttpProxyError("没有远程服务可以被使用")
	}
	//todo 负载均衡...  目前是随机
	addr, err := cs.balance()
	if err != nil {
		return nil, err
	}
	addr = addr + cs.CAPi
	if cs.CPath != "" {
		addr = addr + "?" + cs.CPath
	}

	cmethod := proxy_http.GetMethod(cs.CMethod)
	if cmethod == "" {
		return nil, cs.GoErr.WithHttpProxyError("代理服务不支持此种请求方式" + string(cmethod))
	}

	pr, err := proxy_http.NewProxyRequest(addr, proxy_http.WithReqData(cs.ReqData))
	if err != nil {
		return nil, cs.GoErr.WrapWithHttpProxyError(err, "http代理服务错误")
	}

	response, err := pr.Send(cmethod)
	if err != nil {
		return nil, cs.GoErr.WrapWithHttpProxyError(err, "http代理服务错误")
	}
	return response, nil
}

func (cs *clientServer) parse(resp *http.Response) (*proxy_http.ProxyResponse, error) {
	return proxy_http.Parse(resp)
}

// 建造者
func (cs *clientServer) Do() (*proxy_http.ProxyResponse, error) {

	if err := cs.parseClientRequest(); err != nil {
		return nil, err
	}
	if err := cs.discoveryServer(); err != nil {
		return nil, err
	}
	resp, err := cs.request()

	if err != nil {
		return nil, err
	}
	return cs.parse(resp)
}
