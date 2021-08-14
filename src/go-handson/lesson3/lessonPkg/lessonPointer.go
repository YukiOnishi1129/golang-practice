package lessonPkg

import "fmt"


func LessonPointer() {
	n := 123
	p := &n // nのポインタを取得 (ポインタ型になる) (変数の前に「&」をつけることで、ポインタを取り出せる)
	// fmt.Println("number:", n)
	// fmt.Println("pointer:", p)
	// fmt.Println("value:", *p) // ポインタ型の値を取り出すときは変数の前に「*」をつける

	/*
	* ポインタ変数は同じ型をポインタ変数の値を代入させることができる
	*/
	
	// p2 := &m
	// fmt.Printf("p value:%d, address;%p\n", *p, p)
	// fmt.Printf("p2 value:%d, address;%p\n", *p2, p2)

	// pb := p
	// p = p2
	// p2 = pb
	// fmt.Printf("p value:%d, address;%p\n", *p, p)
	// fmt.Printf("p2 value:%d, address;%p\n", *p2, p2)

	/*
	* ポインタ型変数のポインタ
	*/
	q := &p // nのポインタ変数(p)のポインタ
	m := 10000 
	p2 := &m // mのポインタ
	q2 := &p2 // mのポインタ変数(p2)のポインタ
	// 「*」記号を使うことで、アドレスがある値を遡って取り出すことができる
	// *q: pの値 (mのポインタのアドレス)
	// **q: mの値 (変数値)
	fmt.Printf("q value:%d, address;%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address;%p\n", **q2, *q2)
	pb := p
	p = p2
	p2 = pb
	fmt.Printf("q value:%d, address;%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address;%p\n", **q2, *q2)
}

func LessonArgPointer() {
	n := 123
	fmt.Printf("value%d.\n", n)
	change1(n)
	fmt.Printf("value%d.\n", n) // 変化なし
	// 引数にポインタ変数を指定する
	change2(&n)
	fmt.Printf("value%d.\n", n) // 変化あり

	// change2ではアドレス値を渡しているので、参照元を変更している
	// change1では変数の値のみを渡しているが、
	// 別のアドレスのものにコピーされてそれを元に加工しているだけなので、
	// 元の変数には影響はない(変更は反映されない。)

}

// 値渡し
func change1(n int) {
	n *= 2
}

// 参照渡し
// 引数にポインタ変数を指定する
func change2(n *int) {
	// *nでポインタ変数nの実際の値に代入している
	*n *= 2
}