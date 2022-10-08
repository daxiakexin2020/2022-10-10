package main

import (
	"bytes"
	"fmt"
)

func main() {
	Handle()
}

type F func(s string) string

type FHandle struct {
	FG []F
}

func (fg *FHandle) run() {
	for _, f := range fg.FG {
		fmt.Println(f("test"))
	}
}

func Handle() {
	var f1 F
	f1 = func(s string) string {
		return s
	}
	res := f1("hello")
	fmt.Println(res)

	f2 := func(s string) string {
		b := bytes.Buffer{}
		b.WriteString("b==")
		b.WriteString(s)
		return b.String()
	}

	fg := make([]F, 0)
	fg = append(fg, f2, f1)

	f := &FHandle{
		FG: fg,
	}
	f.run()
}
