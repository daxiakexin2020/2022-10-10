package main

import "30thrid_parth/v2/server"

func main() {
	test01()
}

func test01() {
	server.FetchBaseDataMapping()
	server.GetToken()
}
