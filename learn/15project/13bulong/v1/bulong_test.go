package v1

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	b := make([]byte, 100)
	b[0] = 2
	fmt.Println("b", b)
}
