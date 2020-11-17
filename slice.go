package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("s1=", arr[2:])
	fmt.Println("s2=", arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	updateSlice(s1)
	updateSlice(s2)
	fmt.Println(s1, s2, arr)

	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
}
