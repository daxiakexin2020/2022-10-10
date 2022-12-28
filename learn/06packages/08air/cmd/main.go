package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	test()
}

func test() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/t", t)
	server := &http.Server{Handler: mux, Addr: ":6003"}
	log.Fatal(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Header(), "Hello World")
}
func t(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello t", w.Header().Get("Accept"))
}
