package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
)

func LessonSwitch() {
	x := hello.Input("type a number")
	fmt.Println(x + "月は、")

	switch n , err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("整数値が得られません。")
		fmt.Println(err)
	case 1, 2, 12:
		fmt.Println("冬です。")
	case 3, 4, 5:
		fmt.Println("春です。")
	case 6, 7, 8:
		fmt.Println("夏です。")
	case 9, 10, 11:
		fmt.Println("秋です。")
	default: 
		fmt.Println("月に値ではありませんよ？")
	}
}

func LessonNoConditionSwitch() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Print(x + "は、")
	} else {
		return
	}

	switch {
	    case n%2 == 0 :
			fmt.Println("偶数です。")
		case n%2 == 1 :
			fmt.Println("奇数です。")
	}
}

func LessonFallthroughSwitch() {
	x := hello.Input("type 1~5")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "までの合計は、")
	} else {
		return
	}

	t := 0

	switch n {
	case 5: 
		t += 5
		fallthrough
	case 4: 
		t += 4
		fallthrough
	case 3: 
		t += 3
		fallthrough
	case 2: 
		t += 2
		fallthrough
	case 1: 
		t++
	default:
		fmt.Println("範囲外です。")
		return
	}
	fmt.Println(t, "です。")
}