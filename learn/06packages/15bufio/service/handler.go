package service

import (
	"bufio"
	"io"
)

func NewReader(reader io.Reader) *bufio.Reader {
	return bufio.NewReader(reader)
}

func NewWriter(writer io.Writer) *bufio.Writer {
	return bufio.NewWriter(writer)
}
