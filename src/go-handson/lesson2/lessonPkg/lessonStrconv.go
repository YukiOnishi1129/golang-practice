package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
)

func LessonStrconv() {
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