package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
	"strings"
)

func LessonArrayBasic() {
	x := hello.Input("input data")
	// 文字列をSplitで配列にして取り出す
	ar := strings.Split(x, " ")
	t := 0

	for i:=0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}

func LessonArrayRange() {
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0

	// range: 配列から順にindexと値を取り出して返す
	// for文と組み合わせて使う
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return
err:
	fmt.Println("ERROR!")

}

func LessonArraySlice() {
	// 配列をsliceで定義
	// 変数 := [] 型 {値1, 値2, ...}
	a := [5]int{10, 20, 30, 40, 50}
	// 配列からsliceを生成
	// 変数 := 配列 [開始 : 終了 ]
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 100
	// a[0]を変更したのに、b[0]も変更されている
	// スライスは配列の参照なので、アドレスは同じ
	fmt.Println(a) // [100 20 30 40 50]
	fmt.Println(b) // [100 20 30]

	b[1] = 200
	fmt.Println(a) // [100 200 30 40 50]
	fmt.Println(b) // [100 200 30]
}

func LessonArrayAppend() {
	a := [3]int{10, 20, 30}
	b := a[0:2]
	fmt.Println(a)
	fmt.Println(b)

	// appendでsliceに値を追加する
	// 変数 = append(スライス, 値1, 値2, ...)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
}

func LessonArrayMap() {
	// 配列やスライスはindexで値を管理しているが
	// 番号ではなく名前で管理したい時はmapを使う
	// var 変数 map[キー型] 値型
	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}
	// m["total"] = m["a"] + m["b"] + m["c"]
	// fmt.Println(m)

	// for 変数1, 変数2 := range マップ変数 { ...繰り返す処理... }

	for k, v := range m {
		fmt.Println(k + ":", v)
	}
}