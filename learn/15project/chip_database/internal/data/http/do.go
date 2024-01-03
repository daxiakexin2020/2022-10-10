package http

import (
	"bytes"
	"chip_database/conf"
	"chip_database/internal/data/http/req"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FileHandleProxy struct {
	base string
}

const PROXY_TIMEOUT = time.Second * 20

func NewFileHandleProxy(conf *conf.FileHandleProxyConfig) *FileHandleProxy {
	return &FileHandleProxy{
		base: conf.Base,
	}
}

func (fhp *FileHandleProxy) Format(data []byte, path string, dest interface{}) (interface{}, error) {
	if err := fhp.send(path, data, dest); err != nil {
		return nil, err
	}
	return dest, nil
}

func (fhp *FileHandleProxy) send(path string, data []byte, dest interface{}, opts ...req.Option) error {

	start := time.Now()
	defer func() {
		s := time.Now().Sub(start).String()
		log.Printf("Proxy Cost Time:::::::::::::::::::：%v\n", s)
	}()

	request, err := req.NewRequest(fhp.generateUrl(path), bytes.NewReader(data), opts...)
	if err != nil {
		return err
	}

	return fhp.do(request, dest)
}

func (fhp *FileHandleProxy) generateUrl(path string) string {
	return fhp.base + path
}

func (fhp *FileHandleProxy) do(req *req.Request, dest interface{}) error {

	log.Printf("========================>【Do Request:%v】\n", *req.Request())

	c := http.Client{Timeout: PROXY_TIMEOUT}
	resp, err := c.Do(req.Request())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("========================>【read data len:%d,http status code :%v】\n", len(data), resp.StatusCode)
	if req.IsShowLog() {
		log.Printf("========================>【read data :%v\n】", string(data))
	}
	if err = json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return nil
}
