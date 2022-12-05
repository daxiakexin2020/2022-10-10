package http

import (
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
	Clinet  *http.Client
	Request *http.Request
}

func NewProxyRequest(url string, reqData map[string]interface{}) (*ProxyRequest, error) {

	if reqData != nil {
		//requestBody := new(bytes.Buffer)
		//json.NewEncoder(requestBody).Encode(reqData) //发送携带json格式的数据
	}
	//formValues := nurl.Values{}
	//formValues.Set("name", "name")
	//formDataStr := formValues.Encode()
	//formDataBytes := []byte(formDataStr)
	//formBytesReader := bytes.NewReader(formDataBytes)

	req, err := http.NewRequest(string(GET_METHOD), url, nil)
	req.Form = nurl.Values{}
	req.PostForm = nurl.Values{}
	if err != nil {
		return nil, err
	}

	req.Header.Set(CONTENT_TYPE, string(URLENCODE_TYPE))

	return &ProxyRequest{
		Clinet:  &http.Client{},
		Request: req,
	}, nil
}

func (pr *ProxyRequest) SetContextType(val application_type) *ProxyRequest {
	return pr.SetHeader(CONTENT_TYPE, string(val))
}

func (pr *ProxyRequest) SetHeader(key string, val string) *ProxyRequest {
	pr.Request.Header.Set(key, val)
	return pr
}

func (pr *ProxyRequest) SetHeaders(heads map[string]string) *ProxyRequest {
	for k, v := range heads {
		pr.SetHeader(k, v)
	}
	return pr
}

func (pr *ProxyRequest) SetForm(key string, val string) *ProxyRequest {
	pr.Request.Form.Set(key, val)
	return pr
}

func (pr *ProxyRequest) SetForms(heads map[string]string) *ProxyRequest {
	for k, v := range heads {
		pr.SetForm(k, v)
	}
	return pr
}

func (pr *ProxyRequest) SetPostForm(key string, val string) *ProxyRequest {
	pr.Request.PostForm.Set(key, val)
	return pr
}

func (pr *ProxyRequest) SetPostForms(heads map[string]string) *ProxyRequest {
	for k, v := range heads {
		pr.SetPostForm(k, v)
	}
	return pr
}

func (pr *ProxyRequest) Close() {
	pr.Request.Body.Close()
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

func (pr *ProxyRequest) do(method method_type) (*http.Response, error) {
	pr.Request.Method = string(method)
	return pr.Clinet.Do(pr.Request)
}
