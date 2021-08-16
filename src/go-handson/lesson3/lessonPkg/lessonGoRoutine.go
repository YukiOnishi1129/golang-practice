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

/*
* チャンネル
*/
func total(n int, c chan int) {
	t := 0
	for i := 0; i <= n; i++ {
		t += i
	}
	// 値を追加する
	// チャンネル　<= 値
	c <- t
}

func LessonChannel() {
	// チャンネル型の変数定義
	// 変数　:= make(chan 型)
	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	x, y, z := <-c, <-c, <-c
	// 値を取り出す
	// 変数 := <-チャンネル
	// 「チャンネルから値を取得する場合、その値が送られてくるまで処置を待つ」性質がある
	// fmt.Println("total", <-c)

	// 注意：必ずGoルーチンを呼び出した順に値が追加されるわけではない。
	fmt.Println(x, y,z)
}

func totalCross(c chan int) {
	// チャンネルは値がまだ用意されていない場合は、送られてくるまで待つ
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

/*
* メインチャンネル側からチャンネルを渡すことはできない
* fatal error: all goroutines are asleep - deadlock!
* Goルーチンによるスレッドを実行した後でないと、チャンネルは使えない
*/
func LessonChannelCross() {
	c := make(chan int)
	// これだとエラー (goルーチン実行前にチャンネルに値を設定しているため)
	// 送る側と受け取る側の双方向で値の準備ができていないとエラー
	// c <- 100
	// go totalCross(c)
	// goルーチン実行後にチャンネルに値を設定すれば正常に動作する
	go totalCross(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
}

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "second:["+<-c+"]")
	}
}

func LessonThreadChannel() {
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println()
}


/*
* 双方向でチャンネルをやり取りする
*/
func totalBidirectional(cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += 1
	}
	cr <-t
}

func LessonBidirectional() {
	cs := make(chan int)
	cr := make(chan int)
	go totalBidirectional(cs, cr)
	cs <- 100
	fmt.Println("total:", <-cr)
}