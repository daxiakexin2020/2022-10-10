package stores

import (
	"bufio"
	"io"
	"os"
)

type store struct {
	file   *os.File
	reader *bufio.Reader
	writer *bufio.Writer
}

func (s *store) read(handle func(buf []byte)) error {
	for {
		line, err := s.reader.ReadSlice('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(line) > 0 {
			handle(line)
		}
	}
	return nil
}

func (s *store) write(data []byte) (int, error) {
	n, err := s.writer.Write(data)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (s *store) close() error {
	return s.file.Close()
}

func (s *store) flush() error {
	return s.writer.Flush()
}
