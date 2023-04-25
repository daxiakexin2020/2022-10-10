package jiuqi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type base struct {
	url string
}

const JiuQi_Domain = "https://sc.360.net/custom/sc/v2/popular/"

func (b *base) SetUrl(url string) {
	b.url = url
}

func (b *base) generateUrl(path string) {
	b.url = JiuQi_Domain + path
}

func (b *base) send(params interface{}, dest interface{}, opts ...option) error {

	log.Printf("base send ruquest [param:%+v], [url:%v]\n", params, b.url)
	requestBody := new(bytes.Buffer)
	if err := json.NewEncoder(requestBody).Encode(params); err != nil {
		return err
	}
	log.Printf("[json NewEncoder,json param:%v]\n", requestBody)

	req, err := NewRequest(b.url, requestBody, opts...)
	if err != nil {
		return err
	}

	log.Printf("[do request send info:%v]\n", req.request())

	resp, err := req.Do()
	if err != nil {
		return err
	}

	return b.handle(resp, dest)
}

func (b *base) handle(resp *http.Response, dest interface{}) error {

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("[handle,get data from thrit parth:%v],[http status code:%v]\n", string(data), resp.StatusCode)
	if err = json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return nil
}
