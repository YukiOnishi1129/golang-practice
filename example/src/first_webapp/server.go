package main // 実行可能なプログラムは必ず「main」のパッケージの中にある必要あり

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// バイナリ実行ファイルにコンパイルされる全てのプログラムは「main」関数が1必要
func main() {
	// ルートURL('/')が呼び出されたとき、「handler」関数を呼び出す
	http.HandleFunc("/", handler)
	// ポート8080でサーバを起動する
	http.ListenAndServe(":8080", nil)
}
