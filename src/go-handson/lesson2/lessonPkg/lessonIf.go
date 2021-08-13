package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
)

func LessonIf() {
	x := hello.Input("type a number")
	// n, err := strconv.Atoi(x)

	// if err != nil {
	// 	fmt.Println("ERROR!")
	// 	return
	// }

	// fmt.Println(x + "は、")
	// if n%2 == 0 {
	// 	fmt.Println("偶数です。")
	// } else {
	// 	fmt.Println("奇数です。")
	// }

	// リファクタリング
	// ショートステートメント付きif
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("整数ではありません。")
	}
}