package main // main関数のpackageと宣言
// 1つのパッケージのみ実行できる

import (
	"fmt"
	"time"
)

// HelloWorld

/*
* main関数 (エントリーポイント)
* main関数の中の記述が実行される
 */
func main() {
	fmt.Println("HelloWorld")
	fmt.Println(time.Now())
}