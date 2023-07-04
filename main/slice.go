package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("len: %d, cap: %d, slice: %v\n", len(slice), cap(slice), slice)

	s2 := make([]int, 2)
	copy(s2, slice)

	slice[1] = 100
	s2[1] = 200

	fmt.Println(slice)
	fmt.Println(s2)
}
