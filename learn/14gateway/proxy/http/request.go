package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	nurl "net/url"
)

type method_type string

type application_type string

const (
	URLENCODE_TYPE  application_type = "application/x-www-form-urlencoded"
	JSON_TYPE       application_type = "application/json"
	FORM_DATA_TYPE  application_type = "multipart/form-data"
	TEXT_PLAIN_TYPE application_type = "text/plain"
)

const (
	GET_METHOD     method_type = "GET"
	POST_METHOD    method_type = "POST"
	PUT_METHOD     method_type = "PUT"
	DELETE_METHOD  method_type = "DELETE"
	HEAD_METHOD    method_type = "HEAD"
	OPTIONS_METHOD method_type = "OPTIONS"
)

const CONTENT_TYPE = "Content-Type"

type ProxyRequest struct {
	clinet          *http.Client
	request         *http.Request
	reqData         map[string]interface{} `json:"req_data"`
	applicationType application_type       `json:"application_type"`
}

type Option func(pr *ProxyRequest)

func WithReqData(reqData map[string]interface{}) Option {
	return func(pr *ProxyRequest) {
		pr.reqData = reqData
	}
}

func WithApplicationType(applicationType application_type) Option {
	return func(pr *ProxyRequest) {
		pr.applicationType = applicationType
	}
}

func makeJsonDatas(reqData map[string]interface{}) io.Reader {
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(reqData)
	return requestBody
}

func makeFormData(reqDatas map[string]interface{}) io.Reader {
	formValues := nurl.Values{}
	for k, v := range reqDatas {
		if rv, ok := v.(string); ok {
			formValues.Set(k, rv)
		}
	}
	formDataBytes := []byte(formValues.Encode())
	return bytes.NewReader(formDataBytes)
}

func (pr *ProxyRequest) makeReqData() io.Reader {
	if pr.reqData == nil {
		return nil
	}
	switch pr.applicationType {
	case URLENCODE_TYPE:
		return makeFormData(pr.reqData)
	case JSON_TYPE:
		return makeJsonDatas(pr.reqData)
	default:
		return nil
	}
}

func (pr *ProxyRequest) apply(opts []Option) {
	for _, opt := range opts {
		opt(pr)
	}
}

func NewProxyRequest(url string, opts ...Option) (*ProxyRequest, error) {

	pr := &ProxyRequest{applicationType: JSON_TYPE}
	pr.apply(opts)
	req, err := http.NewRequest(string(GET_METHOD), url, pr.makeReqData())
	if err != nil {
		return nil, err
	}
	pr.clinet = &http.Client{}
	pr.request = req
	pr.SetContextType(pr.applicationType)
	return pr, nil
}

func (pr *ProxyRequest) SetContextType(val application_type) *ProxyRequest {
	return pr.SetHeader(CONTENT_TYPE, string(val))
}

func (pr *ProxyRequest) SetHeader(key string, val string) *ProxyRequest {
	pr.request.Header.Set(key, val)
	return pr
}

func (pr *ProxyRequest) SetHeaders(heads map[string]string) *ProxyRequest {
	for k, v := range heads {
		pr.SetHeader(k, v)
	}
	return pr
}

func (pr *ProxyRequest) Close() error {
	if pr.request.Body == nil {
		return nil
	}
	return pr.request.Body.Close()
}

func (pr *ProxyRequest) Get() (*http.Response, error) {
	return pr.do(GET_METHOD)
}

func (pr *ProxyRequest) POST() (*http.Response, error) {
	return pr.do(POST_METHOD)
}

func (pr *ProxyRequest) PUT() (*http.Response, error) {
	return pr.do(PUT_METHOD)
}
func (pr *ProxyRequest) DELTE() (*http.Response, error) {
	return pr.do(DELETE_METHOD)
}
func (pr *ProxyRequest) HEAD() (*http.Response, error) {
	return pr.do(HEAD_METHOD)
}
func (pr *ProxyRequest) OPTIONS() (*http.Response, error) {
	return pr.do(OPTIONS_METHOD)
}

func (pr *ProxyRequest) Send(method method_type) (*http.Response, error) {
	return pr.do(method)
}

func (pr *ProxyRequest) do(method method_type) (*http.Response, error) {
	pr.request.Method = string(method)
	return pr.clinet.Do(pr.request)
}
