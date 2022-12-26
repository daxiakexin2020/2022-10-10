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
	reqData         interface{}      `json:"req_data"`
	applicationType application_type `json:"application_type"`
}

type Option func(pr *ProxyRequest)

func WithReqData(reqData interface{}) Option {
	return func(pr *ProxyRequest) {
		pr.reqData = reqData
	}
}

func WithApplicationType(applicationType application_type) Option {
	return func(pr *ProxyRequest) {
		pr.applicationType = applicationType
	}
}

func makeJsonData(reqData []byte) io.Reader {
	requestBody := new(bytes.Buffer)
	data := make(map[string]interface{})
	json.Unmarshal(reqData, &data)
	json.NewEncoder(requestBody).Encode(data)
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
		v, ok := pr.reqData.(map[string]interface{})
		if ok {
			return makeFormData(v)
		}
		return nil
	case FORM_DATA_TYPE:
		v, ok := pr.reqData.(map[string]interface{})
		if ok {
			return makeFormData(v)
		}
		return nil
	case JSON_TYPE:
		v, ok := pr.reqData.([]byte)
		if ok {
			return makeJsonData(v)
		}
		return nil
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

	pr := &ProxyRequest{}
	pr.apply(opts)
	pr.applicationType = FORM_DATA_TYPE
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
	if pr.applicationType == "" {
		pr.applicationType = FORM_DATA_TYPE
	}
	return pr.do(POST_METHOD)
}

func (pr *ProxyRequest) PUT() (*http.Response, error) {
	if pr.applicationType == "" {
		pr.applicationType = FORM_DATA_TYPE
	}
	return pr.do(PUT_METHOD)
}
func (pr *ProxyRequest) DELTE() (*http.Response, error) {
	if pr.applicationType == "" {
		pr.applicationType = FORM_DATA_TYPE
	}
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
	//todo
	//body := &bytes.Buffer{}
	//writer := multipart.NewWriter(body)
	//pr.request.Header.Del("Content-Type")
	//pr.request.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	return pr.clinet.Do(pr.request)
}

func GetMethod(method string) method_type {
	switch method_type(method) {
	case GET_METHOD:
		return GET_METHOD
	case POST_METHOD:
		return POST_METHOD
	case PUT_METHOD:
		return PUT_METHOD
	case HEAD_METHOD:
		return HEAD_METHOD
	case DELETE_METHOD:
		return DELETE_METHOD
	case OPTIONS_METHOD:
		return OPTIONS_METHOD
	default:
		return ""
	}
}
