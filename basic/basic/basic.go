package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kkk"

var (
	bb = 4
	cc = false
)

// bb := true 在外面不能使用这种定义变量

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	var a, b, s = 3, 4, "def"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

/**
强制类型转换
*/
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	println(aa, bb, cc, ss)
	euler()
	triangle()
	consts()
	enums()
}
