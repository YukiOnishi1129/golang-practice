package main

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
	// 「src」直下から見たパスを記述する
)

func main() {
	// 数値型
	a,b,c := 100, 200, 300
	fmt.Print("total:")
	fmt.Println(a+b+c)

	// 文字列　→ 数値の型変換
	// strconv
	//  x := hello.

	// 「パッケージ名.メソッド名」で記述
	x := hello.Input("type your name")

	// strconv.Atoi (文字列 → 数値)
	n, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("ERROR!")
		return
	}

	p := float64(n)
	fmt.Println(int(p * 1.1))
}