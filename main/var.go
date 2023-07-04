package main

import "fmt"

func foo() int {
	return 1
}

func main() {
	var a int = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := foo()
	fmt.Println(b)
}
