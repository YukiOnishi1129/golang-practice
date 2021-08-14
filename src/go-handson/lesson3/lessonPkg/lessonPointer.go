package lessonPkg

import "fmt"


func LessonPointer() {
	n := 123
	p := &n // nのポインタを取得 (ポインタ型になる) (変数の前に「&」をつけることで、ポインタを取り出せる)
	fmt.Println("number:", n)
	fmt.Println("pointer:", p)
	fmt.Println("value:", *p) // ポインタ型の値を取り出すときは変数の前に「*」をつける
}