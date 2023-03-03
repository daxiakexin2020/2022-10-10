package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", ":9112")
	if err != nil {
		panic(err)
	}
	var replay string
	err = client.Call("Server.Fetch", "1", &replay)
	if err != nil {
		panic(err)
	}
	fmt.Println("replay:", replay)
}
