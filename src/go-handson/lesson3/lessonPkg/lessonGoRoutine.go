package lessonPkg

import (
	"fmt"
	"strconv"
	"time"
)

func helloRoutine(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func LessonRoutine() {
	go helloRoutine("hello", 50)
	helloRoutine("bye!", 100)
}


/*
* 複数スレッドで変数を共有する
* helloとmainという二つの関数を並行処理させる
* お互いmsgという変数を共有する
*/
func LessonMultipleThread() {
	msg := "start"
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			// scope外の変数msgを編集
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	main := func(n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			// scope外の変数msgを編集
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}

	go hello(60)
	main(100)
}