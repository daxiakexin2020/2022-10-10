package server

import (
	"30thrid_parth/v2/model"
	"30thrid_parth/v2/proxy/jiuqi"
	"log"
)

type BaseDataMappingParam struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func FetchBaseDataMapping() {

	proxy := jiuqi.NewBaseDataMapping(model.NewBaseDataMapping())
	params := &BaseDataMappingParam{
		Code: "10",
		Msg:  "no",
		Data: map[string]interface{}{"id": "1", "tableName": "test_teable_name"},
	}
	destModel, err := proxy.Send(params)
	log.Printf("**************************result**************************: %+v,err:%v", destModel, err)
}
