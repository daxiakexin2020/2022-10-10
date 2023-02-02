package user

import (
	"12lucky_draw/helper"
	"fmt"
	"log"
	"testing"
)

func TestUuid(t *testing.T) {
	id := helper.Uuid()
	fmt.Println("uuid", id)
}

func TestAdd(t *testing.T) {
	username := "zhang zhou"
	user, err := Add(username, V01)
	if err != nil {
		log.Panic("add user err", err)
	}
	fmt.Printf("User=%p\n", &user)
	user.Username = "lisi"
	for _, u := range getGDB().users {
		fmt.Printf("Users=%p\n", u)
		fmt.Printf("Users name=%v\n", u.Username)
	}
}
