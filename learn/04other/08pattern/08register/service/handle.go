package service

import (
	"fmt"
	"sync"
)

type Handler func() error

type proxyRegiseter struct {
	proxy  *regiseter
	errors []error
}

type regiseter struct {
	Handlers []Handler
}

var (
	pRegister      *proxyRegiseter
	singleRegister sync.Once
)

func newProxyRegister() *proxyRegiseter {
	if pRegister == nil {
		singleRegister.Do(func() {
			proxy := &regiseter{
				Handlers: make([]Handler, 0),
			}
			errors := make([]error, 0)
			pRegister = &proxyRegiseter{
				proxy:  proxy,
				errors: errors,
			}
		})
	}
	return pRegister
}

func GetProxyRegiseter() *proxyRegiseter {
	return newProxyRegister()
}

func (pr *proxyRegiseter) Register(handler ...Handler) {
	pr.proxy.Handlers = append(pr.proxy.Handlers, handler...)
}

func (pr *proxyRegiseter) Run() {
	for _, handler := range pr.proxy.Handlers {
		pr.setError(handler())
	}
}

func (ps *proxyRegiseter) setError(err error) {
	ps.errors = append(ps.errors, err)
}

func (pr *proxyRegiseter) CheckErr() error {
	for _, err := range pr.errors {
		if err != nil {
			fmt.Printf("register err=%v\n", err)
			return err
		}
	}
	return nil
}
