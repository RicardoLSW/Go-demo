package main

import (
	"fmt"
	"go-demo/queue"
)

func main() {
	q := queue.Queue{"a"}
	q.Push(2)
	q.Push("c")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
