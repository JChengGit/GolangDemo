package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const File = "token.json"

type AppToken struct {
	Token   string
	Expired int64
}

func read() {
	f, err := os.Open(File)
	if err != nil {
		return
	}
	defer f.Close()

	token := AppToken{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&token)
}

func write() {
	now := time.Now().Unix()
	token := AppToken{"token-token-token-token", now}
	f, err := os.Create(File)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(&token)
	if err != nil {
		fmt.Println("Encode Error")
	}
}

func main() {
	now := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		read()
		//write()
	}
	then := time.Now().UnixMilli()
	fmt.Printf("Time cost: %d \n", then-now)
}
