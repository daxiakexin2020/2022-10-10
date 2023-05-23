package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	n := 13
	one := countDigitOne(n)
	log.Printf("one total:%d\n", one)
}

func countDigitOne(n int) int {
	flag := "1"
	var total int
	for i := n; i >= 0; i-- {
		old := strconv.Itoa(i)
		new := strings.Replace(old, flag, "", -1)
		total += len(old) - len(new)
	}
	return total
}
