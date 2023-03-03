package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type Consumer struct {
}

func (c *Consumer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlpath := html.EscapeString(r.URL.Path)
	switch urlpath {
	case "/":
		HelloHandler(w, r)
	case "/test":
		TestHandler(w, r)
	default:
		fmt.Fprintf(w, "默认值, %q", html.EscapeString(r.URL.Path))
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ZZ", "ok")

	type response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	//此种方式，会自动向w接口写入数据，会调用w接口中Write() 方法，将编码后的json数据自动写入
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&response{Code: 100, Msg: "ok"}); err != nil {
		http.Error(w, err.Error(), 500)
	}

	//w.Write([]byte("hello writer\n"))
	//fmt.Fprintf(w, "HelloHandler, %q", html.EscapeString(r.URL.Path))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestHandler, %q", html.EscapeString(r.URL.Path))
}

func main() {

	/**
	一个client是可以复用的
	*/
	//直接发送http请求,  Get方法底层使用了默认的Client配置
	url := "https://sc.360.net"
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	body := resp.Body
	defer body.Close()
	b := make([]byte, 1024)
	body.Read(b)
	fmt.Println("b:", string(b))
	fmt.Println("***********************************")

	//复杂一点的操作 要管理HTTP客户端的头域、重定向策略和其他设置，自己创建一个Client，进行设置
	client := http.Client{
		Timeout: time.Second * 3,
	}
	resp2, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	b2 := make([]byte, 1024)
	resp2.Body.Read(b2)
	fmt.Println("b2:", string(b2))

	fmt.Println("***********************************")

	//还可以使用NewRequest,进行更复杂的操作，比如设置header头等，然后使用client发送请求，最终都会落在client上
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("test", "test")
	resp3, err := client.Do(req)
	defer resp3.Body.Close()
	b3 := make([]byte, 1024)
	resp3.Body.Read(b3)
	fmt.Println("b3:", string(b3))
	fmt.Println("***********************************")

	//要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: nil},
		DisableCompression: true,
	}
	client2 := &http.Client{Transport: tr}
	resp4, err := client2.Get(url)
	defer resp4.Body.Close()
	b4 := make([]byte, 1024)
	resp4.Body.Read(b4)
	fmt.Println("b4:", string(b3))
	fmt.Println("***********************************")

	//自定义，自己可以由此写框架，在serveHttp中进行路由的拆解
	c := &Consumer{}
	s := &http.Server{
		Addr:    ":9111",
		Handler: c}
	s.ListenAndServe()

	//直接注册多个路由
	//底层，还是使用server，http.ListenAndServe 是一个门面
	//func ListenAndServe(addr string, handler Handler) error {
	//	todo
	//	server := &Server{Addr: addr, Handler: handler}
	//	return server.ListenAndServe()
	//}
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/test", TestHandler)
	http.ListenAndServe(":9112", nil)

}
