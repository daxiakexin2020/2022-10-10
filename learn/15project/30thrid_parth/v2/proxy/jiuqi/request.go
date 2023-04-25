package jiuqi

import (
	"io"
	"net/http"
)

type request struct {
	req    *http.Request
	client *http.Client
	err    error
}

type method string

type option func(r *request)

const (
	DefaultMethod      method = "POST"
	DefaultContentType        = "application/json"
	ContentType               = "Content-Type"
)

const (
	GET  method = "GET"
	POST        = "POST"
)

func WithMethod(method method) option {
	return func(r *request) {
		r.SetMethod(method)
	}
}

func (r *request) apply(opts ...option) {
	r.SetHeader(ContentType, DefaultContentType)
	for _, opt := range opts {
		opt(r)
	}
}

func NewRequest(url string, data io.Reader, opts ...option) (*request, error) {
	req, err := http.NewRequest(string(DefaultMethod), url, data)
	if err != nil {
		return nil, err
	}
	r := &request{req: req, client: http.DefaultClient, err: nil}
	r.apply(opts...)
	return r, nil
}

func (r *request) SetHeader(key string, val string) {
	r.req.Header.Set(key, val)
}

func (r *request) SetMethod(method method) {
	r.req.Method = string(method)
}

func (r *request) Do() (*http.Response, error) {
	return r.client.Do(r.req)
}

func (r *request) SetClient(c *http.Client) {
	r.client = c
}

func (r *request) request() *http.Request {
	return r.req
}

func (r *request) SetError(err error) {
	if r.err == nil {
		r.err = err
	}
}

func (r *request) GetError() error {
	return r.err
}
