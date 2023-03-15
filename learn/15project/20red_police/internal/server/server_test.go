package server

import (
	"20red_police/protocol"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

var conn net.Conn

var (
	RegisterRes   = make(chan interface{}, 1)
	LoginRes      = make(chan interface{}, 1)
	CreateRoomRes = make(chan interface{}, 1)
	CreatePMapRes = make(chan interface{}, 1)
	JoinRoomRes   = make(chan interface{}, 10)
)

type Common struct {
	ServiceMethod string `json:"service_method"`
}

type RegisterRequest struct {
	Common
	MetaData protocol.RegisterRequest `json:"meta_data"`
}

type RegisterResponse struct {
	Data protocol.RegisterResponse `json:"data"`
	Err  string                    `json:"err"`
}

type LoginRequest struct {
	Common
	MetaData protocol.LoginRequest `json:"meta_data"`
}

type LoginResponse struct {
	Data protocol.LoginResponse `json:"data"`
	Err  string                 `json:"err"`
}

type CreateRoomResponse struct {
	Data protocol.CreateRoomResponse `json:"data"`
	Err  string                      `json:"err"`
}

type CreateRoomRequest struct {
	Common
	MetaData protocol.CreateRoomRequest `json:"meta_data"`
}

type CreatePMapRequest struct {
	Common
	MetaData protocol.CreatePMapRequest `json:"meta_data"`
}

type CreatePMapResponse struct {
	Data protocol.CreatePMapResponse `json:"data"`
	Err  string                      `json:"err"`
}

type JoinRoomRequest struct {
	Common
	MetaData protocol.JoinRoomRequest `json:"meta_data"`
}

type JoinRoomResponse struct {
	Data protocol.JoinRoomResponse `json:"data"`
	Err  string                    `json:"err"`
}

func TestRegisterLoginCreateRoom(t *testing.T) {
	mapping := map[string]func(t *testing.T){
		"TestRegister":   TestRegister,
		"TestLogin":      TestLogin,
		"TestCreatePMap": TestCreatePMap,
		"TestCreateRoom": TestCreateRoom,
	}
	dial, err := client()
	if err != nil {
		t.Error(err)
	}
	conn = dial
	defer conn.Close()
	for name, f := range mapping {
		t.Run(name, f)
	}
}

func TestRegister(t *testing.T) {
	req := &RegisterRequest{}
	res := &RegisterResponse{}
	req.Common.ServiceMethod = "Server.Register"
	req.MetaData.Name = "zz08"
	req.MetaData.Pwd = "123"
	req.MetaData.RePwd = "123"
	req.MetaData.Phone = "15248164198"
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("register req err:", err)
	}
	fmt.Println("register req:", string(data))
	conn.Write(data)
	read(conn, RegisterRes, res)
}

func TestLogin(t *testing.T) {
	//registerRes := <-RegisterRes
	//rs := registerRes.(*RegisterResponseData)
	//fmt.Println("register res:", rs)
	req := &LoginRequest{}
	req.Common.ServiceMethod = "Server.Login"
	req.MetaData.Name = "zz08"
	req.MetaData.Pwd = "123"
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("login req marshal err:", err)
	}
	conn.Write(data)
	res := &LoginResponse{}
	read(conn, LoginRes, res)
}

func TestCreateRoom(t *testing.T) {
	loginRes := <-LoginRes
	ls := loginRes.(*LoginResponse)
	fmt.Println("login res:", ls)
	res := &CreateRoomResponse{}
	req := CreateRoomRequest{}
	req.Common.ServiceMethod = "Server.CreateRoom"
	//req.MetaData= ls.Data.Cookie
	//req.MetaData.BName = ls.Data.BName
	req.MetaData.RoomName = "冰天雪地"
	req.MetaData.PMapID = "fd7fd0c8-c246-11ed-2a8a-379c3320538f"
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("create room req:", string(data))
	conn.Write(data)
	read(conn, CreateRoomRes, &res)
	fmt.Println("create room res:", res)
	LoginRes <- loginRes
}

func TestCreatePMap(t *testing.T) {
	loginRes := <-LoginRes
	ls := loginRes.(*LoginResponse)
	fmt.Println("login res:", ls)
	res := &CreatePMapResponse{}
	req := CreatePMapRequest{}
	req.Common.ServiceMethod = "Server.CreatePMap"
	//req.MetaData.Base.Cookie = ls.Data.Cookie
	//req.MetaData.Base.BName = ls.Data.BName
	req.MetaData.Name = "冰天雪地"
	req.MetaData.Count = 8
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("create pmap req:", string(data))
	conn.Write(data)
	read(conn, CreatePMapRes, &res)
	fmt.Println("create pmap res:", res)
	LoginRes <- loginRes
}

func client() (net.Conn, error) {
	dial, err := net.Dial("tcp4", ":9115")
	if err != nil {
		fmt.Println("dial err:", err)
		return nil, err
	}
	return dial, nil
}

func TestJoinRoom(t *testing.T) {
	limit := 10
	for i := 0; i < limit; i++ {
		go func(i int) {
			dial, _ := net.Dial("tcp4", ":9115")
			req := &JoinRoomRequest{}
			res := &JoinRoomResponse{}
			req.Common.ServiceMethod = "Server.JoinRoom"
			req.MetaData.Username = "zz_" + strconv.Itoa(i)
			req.MetaData.RoomID = "123"
			data, err := json.Marshal(req)
			if err != nil {
				fmt.Println("join room req err:", err)
			}
			fmt.Println("join room req:", string(data))
			dial.Write(data)
			read(dial, JoinRoomRes, res)
		}(i)
	}
	time.Sleep(time.Second * 3)
	for data := range JoinRoomRes {
		fmt.Println(data)
	}
}

func read(conn net.Conn, resc chan interface{}, req interface{}) {
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&req); err != nil {
		fmt.Println("Err:", err)
	}
	resc <- req
}
