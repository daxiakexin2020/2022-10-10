package v1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"v1/random"
)

type Server struct {
	r    *random.Rand
	path string
}

func NewServer(path string) *Server {
	return &Server{
		r:    random.NewRand(),
		path: path,
	}
}

const (
	BIG_FILE  = "big_file.txt"
	BIG_LIMIT = 1000000
)

func (s *Server) InitializeBigFile() {
	file, err := os.OpenFile(s.path+"/"+BIG_FILE, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for j := 0; j < BIG_LIMIT; j++ {
		count := s.r.MakeCount(int32(j))
		var str strings.Builder
		var i int32
		for i = 1; i <= count; i++ {
			str.WriteString(strconv.FormatInt(int64(j), 10))
			str.WriteString("\n")
		}
		_, err := file.WriteString(str.String())
		if err != nil {
			fmt.Println("	【file write error 】: ", err)
		}
	}
}
