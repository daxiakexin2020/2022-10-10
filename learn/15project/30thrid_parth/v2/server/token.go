package server

import (
	"30thrid_parth/v2/model"
	"30thrid_parth/v2/proxy/jiuqi"
	"log"
)

type TokenParam struct {
	Openid string `json:"openid"`
}

func GetToken() {
	proxy := jiuqi.NewToken(model.NewToken())
	destModel, err := proxy.Send(&TokenParam{Openid: "123456"})
	log.Printf("**************************result************************** %+v\n,err:%v", destModel, err)
}
