package main

import "log"

func main() {
	Handle()
}

func Handle() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover=================")
			log.Println(err)
		}
	}()
	panic("test===")
}
