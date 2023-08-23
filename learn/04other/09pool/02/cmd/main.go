package main

import (
	"02/service"
)

func main() {
	handle()
}

func handle() {
	service.Test()
}

type Base struct {
	Uid       string `json:"uid"`
	CompanyId string `json:"company_id"`
}

type FetchCompanyTreeRequst struct {
	Base
}

//func test() {
//	a := &FetchCompanyTreeRequst{}
//}
