package network

import (
	"encoding/json"
	"io"
	"log"
)

type Response struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

func sendResponse(data interface{}, err error, writer io.WriteCloser) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	res := &Response{Data: data, Err: errMsg}
	if err = json.NewEncoder(writer).Encode(res); err != nil {
		log.Println("send response err:", err)
	}
}
