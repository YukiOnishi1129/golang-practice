package main

import (
	"fmt"
	"go-handson/package/hello"
	// 「src」直下から見たパスを記述する
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello," + name + "!!")
}