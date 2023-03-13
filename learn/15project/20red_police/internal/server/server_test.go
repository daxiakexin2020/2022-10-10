package server

import (
	"20red_police/config"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"testing"
)

var conn net.Conn

func TestRegisterLoginCreateRoom(t *testing.T) {
	mapping := map[string]func(t *testing.T){
		"TestRegister":   TestRegister,
		"TestLogin":      TestLogin,
		"TestCreateRoom": TestCreateRoom,
	}
	dial, err := client()
	if err != nil {
		t.Error(err)
	}
	conn = dial
	for name, f := range mapping {
		t.Run(name, f)
	}
}

func TestRegister(t *testing.T) {
	fmt.Println("TestRegister start")
	var wg sync.WaitGroup
	data := []byte("{\"service_method\":\"Server.Register\",\"meta_data\":{\"name\":\"zz08\",\"pwd\":\"123\",\"repwd\":\"123\",\"phone\":\"45\"}}")
	conn.Write(data)
	wg.Add(1)
	go read(conn, &wg)
	wg.Wait()
	fmt.Println("register over")
}

func TestLogin(t *testing.T) {
	fmt.Println("TestLogin start")
	var wg sync.WaitGroup
	data := []byte("{\"service_method\":\"Server.Login\",\"meta_data\":{\"name\":\"zz08\",\"pwd\":\"123\"}}")
	conn.Write(data)
	wg.Add(1)
	go read(conn, &wg)
	wg.Wait()
	fmt.Println("login over")
}
func TestCreateRoom(t *testing.T) {
	fmt.Println("TestCreateRoom start")
	var wg sync.WaitGroup
	d := "{\"service_method\":\"Server.CreateRoom\",\"meta_data\":{\"base\":{\"cookie\":\"1\",\"bname\":\"zz08\"},\"room_name\":\"r01\",\"pmap_id\":\"1\"}}"
	data := []byte(d)
	conn.Write(data)
	wg.Add(1)
	go read(conn, &wg)
	wg.Wait()
	fmt.Println("TestCreateRoom over")
}

var successCount int32
var errCount int32

func client() (net.Conn, error) {
	dial, err := net.Dial("tcp4", ":9115")
	if err != nil {
		fmt.Println("dial err:", err)
		return nil, err
	}
	return dial, nil
}

func read(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
L:
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil && err != io.EOF {
			fmt.Println("read err:", err)
			atomic.AddInt32(&errCount, 1)
			break L
		}
		if n != 0 {
			log.Println("recive server data:", string(b))
			atomic.AddInt32(&successCount, 1)
			break L
		}
	}
}

func concurrenceRun(data []byte) {
	var wg sync.WaitGroup
	limit := 2
	addr := fmt.Sprintf("%s:%d", config.GetGrpcServerConfig().Addr, config.GetGrpcServerConfig().Port)
	fmt.Println(addr)
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dial, err := net.Dial("tcp4", addr)
			if err != nil {
				fmt.Println("dial err:", err)
				return
			}
			_, err = dial.Write([]byte("{\"service_method\":\"Server.Register\",\"meta_data\":{\"name\":\"zz08\",\"pwd\":\"123\",\"repwd\":\"123\",\"phone\":\"45\"}}"))
			if err != nil {
				fmt.Println("dial write err:", err)
			}
			wg.Add(1)
			go read(dial, &wg)
		}()
	}
	wg.Wait()
	fmt.Printf("over,count=%d,successCount=%d,errCount=%d\n", limit, successCount, errCount)
}
