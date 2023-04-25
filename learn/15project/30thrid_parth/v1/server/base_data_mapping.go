package server

import (
	"30thrid_parth/v1/proxy/jiuqi"
	"fmt"
)

type BaseDataMappingTest struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func FetchBaseDataMapping() {
	pdata := jiuqi.NewBaseDataMapping()
	data := map[string]interface{}{
		"Id":        "1",
		"TableName": "test_teable_name",
	}
	b := &BaseDataMappingTest{
		Code: "10",
		Msg:  "no",
		Data: data,
	}
	mapping, err := pdata.Send(b)
	fmt.Println("send result :::::", mapping, err)
}
