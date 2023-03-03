package main

import (
	"fmt"
	"net/mail"
)

func main() {
	email := "578975595@qq.com"
	address, err := mail.ParseAddress(email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parseAddress address=%s,name=%s", address.Address, address.Name)
}
