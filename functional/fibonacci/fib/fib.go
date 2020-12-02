package fib

import (
	"fmt"
	"io"
	"strings"
)

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	str := fmt.Sprintf("%d\n", next)
	return strings.NewReader(str).Read(p)
}
