package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := input("type your name")
	fmt.Println("Hello," + name + "!!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	// 入力待ちになるので、入力してEnterをクリック
	scanner.Scan()
	return scanner.Text()
}