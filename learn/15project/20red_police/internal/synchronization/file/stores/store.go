package stores

import (
	"encoding/json"
	"log"
	"os"
)

type store struct {
	file *os.File
}

func (s *store) read(buf []byte) (int, error) {
	n, err := s.file.Read(buf)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (s *store) write(data interface{}) (int, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Println("file marshal err:", err)
		return 0, err
	}
	marshal = append(marshal, '\n')
	n, err := s.file.Write(marshal)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (s *store) close() error {
	return s.file.Close()
}
