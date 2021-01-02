package main // packageは1つのみ

import "fmt"

// 暗黙的な定義は関数内でしか定義できない
// i5 := 300

// 明示的な定義は関数外でもできる
var i6 int = 300

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	fmt.Println(i6)

	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	// まとめて定義
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数の型だけ
	// 初期値が入る
	var i3 int    // 初期値 0
	var s3 string // 初期値 ""
	fmt.Println(i3, s3)

	// 後ほど代入できる
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	// 関数内でしか定義できない
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 再定義できない
	// i4 := 450

	// 異なる型の値を再代入できない
	// i4 = "Hello"

	// 関数を呼び出し
	outer()

	// 使われていない変数があるとエラーになる
	var s5 string = "Not use"
}
