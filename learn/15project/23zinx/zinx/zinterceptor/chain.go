/**
 * @author uuxia
 * @date 15:56 2023/3/10
 * @description 责任链模式
 **/

package zinterceptor

import "github.com/aceld/zinx/ziface"

type Chain struct {
	req          ziface.IcReq
	position     int
	interceptors []ziface.IInterceptor
}

func (c *Chain) Request() ziface.IcReq {
	return c.req
}

func (c *Chain) Proceed(request ziface.IcReq) ziface.IcResp {
	if c.position < len(c.interceptors) {
		//这是下一个需要执行的拦截器信息 c.position+1
		chain := NewChain(c.interceptors, c.position+1, request)

		//这是当前需要执行的拦截器  c.position
		interceptor := c.interceptors[c.position]
		response := interceptor.Intercept(chain)
		return response
	}
	return request
}

func NewChain(list []ziface.IInterceptor, pos int, req ziface.IcReq) ziface.IChain {
	return &Chain{
		req:          req,
		position:     pos,
		interceptors: list,
	}
}
