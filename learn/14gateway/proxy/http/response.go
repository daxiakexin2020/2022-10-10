package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type ProxyResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Parse(resp *http.Response) (*ProxyResponse, error) {

	var sbuf []byte
	pr := &ProxyResponse{}
	step := 1024

	for {
		buf := make([]byte, step)
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		sbuf = append(sbuf, buf[:n]...)

		if n < step {
			break
		}
	}
	if err := json.Unmarshal(sbuf, pr); err != nil {
		return nil, err
	}
	return pr, nil
}
