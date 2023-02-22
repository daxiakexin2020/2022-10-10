package main

import (
	"13table_driver/server"
	"fmt"
)

func main() {
	test()
}

func test() {
	s, err := server.NewServe(1)
	if err != nil {
		panic(err)
	}
	str := s.Srv.Cat()
	fmt.Printf("res=%s\n", str)
	s.Srv.Set("111")

	s2, err := server.NewServe(1)
	if err != nil {
		panic(err)
	}
	str2 := s2.Srv.Cat()
	fmt.Printf("res2=%s\n", str2)

	//rs1 := reflect.ValueOf(s.Srv)
	//rs2 := reflect.ValueOf(s2.Srv)
	//
	//if reflect.DeepEqual(rs1, rs2) {
	//	fmt.Println("相等")
	//} else {
	//	fmt.Println("不相等")
	//}
}
