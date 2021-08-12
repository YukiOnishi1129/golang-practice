package main

import "fmt"

// 変数

// 暗黙的な定義は関数内でしか定義できない
// i5 := 500

var i5 = 500

func outer() {
	var s4 string = "outer"
	fmt.Println((s4))
}

func main () {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println((s))

	var t,f bool = true, false
	fmt.Println(t)
	fmt.Println(f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 値を入れないと初期値が入る
	var i3 int // 0
	var s3 string // ""空文字
	fmt.Println(i3, s3)

	i3 = 300 
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義 (関数内でしか定義できない)
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 暗黙的な定義は再定義できない
	// i4 := 400
	// fmt.Println(i4)

	// 異なる型はエラーになる
	// ./main.go:50:5: cannot use "Hello" (type untyped string) as type int in assignment
	// i4 = "Hello"
	// fmt.Println(i4)


	fmt.Println(i5)

	// 関数街に定義された関数はmain関数の中で呼び出す
	outer()

	// Goは定義した変数を必ず使わないといけない
	// 使わないとエラーになる
	// ./main.go:71:6: s5 declared but not used
	// var s5 string = "Not use"
}