package jiuqi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Base struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const JiuQi_Domain = "https://sc.360.net/custom/sc/v2/popular/"

func (b *Base) generateUrl(path string) string {
	return JiuQi_Domain + path
}

func (b *Base) send(url string, param interface{}, dest interface{}, opts ...option) error {

	log.Printf("base send ruquest [param:%+v], [url:%v]\n", param, url)
	requestBody := new(bytes.Buffer)
	if err := json.NewEncoder(requestBody).Encode(param); err != nil {
		return err
	}
	log.Printf("[json NewEncoder,json param:%v]\n", requestBody)

	request, err := NewRequest(url, requestBody, opts...)
	if err != nil {
		return err
	}

	return b.do(request, dest)
}

func (b *Base) do(req *Request, dest interface{}) error {

	resp, err := http.DefaultClient.Do(req.request())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("[read data:%v],[http status code:%v]\n", string(data), resp.StatusCode)
	if err = json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return nil
}

func (b *Base) isSuccess() bool {
	return b.Code == "0"
}
