package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100

	fmt.Println(i + 50)

	var i2 int64 = 200
	fmt.Println(i2)
	// 同じintでも型が違うとエラーになる
	// invalid operation: i + i2 (mismatched types int and int64)
	// fmt.Println(i + i2)

	// 型を調べる"%T\n"
	fmt.Printf("%T\n",i2) // int64


	fmt.Printf("%T\n", int32(i2))
	// 型変換
	// int32(i2) int64からint32へ
}