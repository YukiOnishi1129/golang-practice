package hello

import (
	"bufio"
	"fmt"
	"os"
)

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	// 入力待ちになるので、入力してEnterをクリック
	scanner.Scan()
	return scanner.Text()
}