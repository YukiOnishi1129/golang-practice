package lessonPkg

import "fmt"

func LessonConst() {
	// 定数
	// 型を固定化すると
	const n1 int = 100
	// これでエラー
	// m := n1 * 1.1

	const n2 = 100
	m := n2 * 1.1
	fmt.Println(m)
}