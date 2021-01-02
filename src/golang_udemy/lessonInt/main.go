package main

import "fmt"

// 型
// 整数型

func main() {
	// intは環境依存
	var i int = 100

	var i2 int64 = 200
	fmt.Println(i + 50)
	// 同じint型でも環境依存のintとint64などは全く別の型になる
	// fmt.Println(i + i2)

	// 型を調べる "%T\n"
	fmt.Printf("%T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}
