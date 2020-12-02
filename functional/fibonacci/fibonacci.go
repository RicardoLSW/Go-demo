package main

import (
	"bufio"
	"fmt"
	"go-demo/functional/fibonacci/fib"
	"io"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printFileContents(f)
}
