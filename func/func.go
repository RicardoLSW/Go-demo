package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
