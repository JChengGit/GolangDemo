package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
}

func (u *User) Call() {
	fmt.Println("user id called.")
}

func main() {
	user := User{1, "foo"}
	user.Call()
	uType := reflect.TypeOf(user)
	fmt.Println(uType.Name())
	fmt.Println(uType.NumField())
	fmt.Println(uType.NumMethod())
}
