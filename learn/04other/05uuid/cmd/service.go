package main

import (
	"fmt"
	"github.com/agext/uuid"
)

func main() {
	Handle()
}

func Handle() {
	res := uuid.NodeId()
	fmt.Print(res)
}
