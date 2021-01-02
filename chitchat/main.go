package main

import (
	"net/http"
)

func main() {

	// マルチプレクサの生成(リクエストをハンドラにリダイレクトするコード)
	mux := http.NewServeMux()
	// FileServer: 指定したディレクトリからファイルを送信するハンドラ
	files := http.FileServer(http.Dir("/public"))
	// StripPrefix: プレフィックスを削除する関数
	// /static/ではじまるリクエストURLに対し、URLから文字列/static/を取り去って、
	// publicを起点として、残った文字列を名前にもつファイルを探す
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFunc: ルートURLをハンドラ関数にリダイレクトさせる関数
	// 第１引数：URL
	// 第２引数：ハンドラ関数
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
