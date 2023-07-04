package main

import "fmt"

//func main() {
//	fmt.Println(strings.Join(os.Args[1:], "->"))
//}

func main() {
	a := 10
	b := 20
	fmt.Println(&a)
	fmt.Println(&b)
}
