package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
)

func LessonFor() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Print("1から"+x + "の合計は、")
	} else {
		return
	}
	t := 0
	// c := 1

	// for c <= n {
	// 	t += c
	// 	c++
	// }
	

	for i := 1; i <= n; i++ {
		t += 1
	}
	fmt.Println(t, "です。")
}

func LessonForContinue() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Print("1から"+x + "の偶数の合計は、")
	} else {
		return
	}

	t := 0
	c := 1

	// for { 繰り返す処理 }
	// breakを組み込む

	for {
		c++
		if c%2 == 1 {
			continue
		}
		if c > n {
			break
		}
		t += c
	}
	fmt.Println(t, "です。")
}