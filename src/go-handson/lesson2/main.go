package main

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
	// 「src」直下から見たパスを記述する
)

func main() {
	lessonInt()
	lessonStrconv()
	lessonConst()
}

func lessonInt() {
	// 数値型
	a,b,c := 100, 200, 300
	fmt.Print("total:")
	fmt.Println(a+b+c)
}

func lessonStrconv() {
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

func lessonConst() {
	// 定数
	// 型を固定化すると
	const n1 int = 100
	// これでエラー
	// m := n1 * 1.1

	const n2 = 100
	m := n2 * 1.1
	fmt.Println(m)
}