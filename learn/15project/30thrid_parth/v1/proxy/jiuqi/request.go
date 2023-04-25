package jiuqi

import (
	"io"
	"net/http"
)

type Request struct {
	req *http.Request
	err error
}

type method string

type option func(r *Request)

const (
	DefaultMethod      method = "GET"
	DefaultContentType        = "application/json"
	ContentType               = "Content-Type"
)

const (
	GET  method = "GET"
	POST        = "POST"
)

func WithMethod(method method) option {
	return func(r *Request) {
		r.SetMethod(method)
	}
}

func (r *Request) apply(opts ...option) {
	r.SetHeader(ContentType, DefaultContentType)
	for _, opt := range opts {
		opt(r)
	}
}

func NewRequest(url string, data io.Reader, opts ...option) (*Request, error) {
	req, err := http.NewRequest(string(DefaultMethod), url, data)
	if err != nil {
		return nil, err
	}
	r := &Request{req: req, err: nil}
	r.apply(opts...)
	return r, nil
}

func (r *Request) SetHeader(key string, val string) {
	r.req.Header.Set(key, val)
}

func (r *Request) SetMethod(method method) {
	r.req.Method = string(method)
}

func (r *Request) SetError(err error) {
	if r.err == nil {
		r.err = err
	}
}

func (r *Request) GetError() error {
	return r.err
}

func (r *Request) request() *http.Request {
	return r.req
}
