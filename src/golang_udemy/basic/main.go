package main

import "fmt"

// 型
// 整数型
func lessonInt() {
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

// 浮動小数点型
func lessonFloat() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義の場合、自動で浮動小数点型になる
	// 環境依存ではない
	fl := 3.2
	fmt.Println(fl)
	// 型を調べる
	fmt.Printf("%T, %T\n", fl64, fl)

	// 明示的にfloat32
	// float32はあまり使わない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	// 0で割ると「+Int」 正の無限大になる
	fmt.Println(pinf)

	ninf := -1.0 / zero
	// 0で割ると「-Int」 負の無限大になる
	fmt.Println(ninf)

	nan := zero / zero
	// 0を0で割ると「NaN」
	fmt.Println(nan)
}

// 論理値型
func lessonBool() {
	var t, f bool = true, false
	fmt.Println(t, f)
}

// 文字列型
func lessonString() {
	// 文字列型はダブルコーテーション
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行の場合、バックコーテーション
	fmt.Println(`test
	test
		test
	`)
	// エスケープシーケンス
	fmt.Println("\"")
	fmt.Println(`"`)

	// string型はbyte型なので、要素として取り出せる。
	// そのまま取り出すと数値で返ってくる
	fmt.Println(s[0])
	// string型の変換すると、文字が表示される
	fmt.Println(string(s[0]))
}

func main() {
	// int型
	// lessonInt()

	// float型
	// lessonFloat()

	// boolean型
	// lessonBool()

	// string型
	lessonString()
}
