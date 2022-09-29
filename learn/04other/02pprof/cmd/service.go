package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
)

type Test struct {
}

func main() {
	service()
}

func service() {
	fmt.Println("====================start====================")
	http.HandleFunc("/test", Test1)
	http.HandleFunc("/test2", Test2)
	http.ListenAndServe(":8083", nil)
	fmt.Print("====================ok====================")
}

func Test1(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintln(os.Stdout, "test")
	r.Write([]byte("test byte"))
}

func Test2(r http.ResponseWriter, w *http.Request) {
	//fmt.Fprintln(os.Stdout, "test2")
	fmt.Println("ddddddddddddddddddddd")
	m := sync.Mutex{}
	for i := 0; i < 300; i++ {
		go func(data int) {
			m.Lock()
			defer m.Unlock()
			fmt.Printf("data=%d\n", data)
			file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666))
			n, err := file.Write([]byte(strconv.Itoa(i) + "\n"))
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			log.Println(n)
			testTime()
		}(i)
	}
}

func testTime() {
	m := make([]int, 0)
	tmpStr := ""
	for i := 0; i < 100000; i++ {
		fmt.Println("test time ok")
		m = append(m, i)
		for j := 0; j < 300; j++ {
			d := i * j
			fmt.Println(d)
			tmpStr += strconv.Itoa(d)
			fmt.Println(tmpStr)
		}
	}
	fmt.Println("test time ok")
}
