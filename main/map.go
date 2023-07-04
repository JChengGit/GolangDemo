package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	if map1 == nil {
		fmt.Println("nil")
	}
	map1["a"] = "1"

	fmt.Println(map1)

	map2 := map[int]string{
		1: "python",
		2: "java",
		3: "golang",
	}
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, "-", value)
	}

}
