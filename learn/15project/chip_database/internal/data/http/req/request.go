package req

import (
	"io"
	"net/http"
)

type Request struct {
	req       *http.Request
	err       error
	isShowLog bool
}

type method string

type Option func(r *Request)

const (
	DefaultMethod method = "POST"
	//DefaultContentType        = "application/json"
	DefaultContentType = "multipart/form-data"
	ContentType        = "Content-Type"
)

const (
	GET  method = "GET"
	POST        = "POST"
)

const (
	AuthorizationKey         = "Authorization"
	AuthorizationBearerToken = "Bearer "
)

func WithMethod(method method) Option {
	return func(r *Request) {
		r.SetMethod(method)
	}
}

func WithIsShowLog(flag bool) Option {
	return func(r *Request) {
		r.isShowLog = flag
	}
}

func (r *Request) apply(opts ...Option) {
	r.SetHeader(ContentType, DefaultContentType)
	for _, opt := range opts {
		opt(r)
	}
}

func NewRequest(url string, data io.Reader, opts ...Option) (*Request, error) {
	req, err := http.NewRequest(string(DefaultMethod), url, data)
	if err != nil {
		return nil, err
	}
	r := &Request{req: req, err: nil, isShowLog: true}

	r.apply(opts...)
	return r, nil
}

func (r *Request) SetBasicAuth(username, password string) {
	r.req.SetBasicAuth(username, password)
}

func (r *Request) SetAuthorizationBearerToken(token string) {
	r.SetHeader(AuthorizationKey, AuthorizationBearerToken+" "+token)
}

func (r *Request) SetHeader(key string, val string) {
	r.req.Header.Set(key, val)
}

func (r *Request) IsShowLog() bool {
	return r.isShowLog
}

func (r *Request) SetMethod(method method) {
	r.req.Method = string(method)
}

func (r *Request) Request() *http.Request {
	return r.req
}
