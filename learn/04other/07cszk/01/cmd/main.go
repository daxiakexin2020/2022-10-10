package main

import (
	"fmt"
	"os"
)

func main() {
	handle()
}

func handle() {
	getenv := os.Getenv("XA_CHANNEL_ID")
	fmt.Println("getenv", getenv)
}
