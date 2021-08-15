package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
)

var myData struct {
	Name string
	Data []int
}

// 構造体の名前の初めは大文字から始まるものにする
// 外部から利用可能なものは、初めは大文字にする
// typeとして型定義
type MyData struct {
	Name string
	Data []int
}

func LessonStruct() {
	// myData.Name = "Taro"
	// myData.Data = []int{10, 20, 30}
	// fmt.Println(myData)
	taro := MyData{"Taro", []int{10, 20, 30}}
	hanako := MyData{
		Name: "Hanako", 
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
}

func LessonStructReference() {
	taro := MyData{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	rev(&taro)
	fmt.Println(taro)
}

// 構造体は参照渡しする
// 値渡しをするとコピーが量産されメモリを圧迫してしまうから
func rev(md * MyData) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) -1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}

func LessonNewStruct() {
	// new: 値を生成し、そのポインタを返す
	// 指定された型の値を生成するためのもの
	// 構造体が持つ変数などは初期化されないので、作成した後に代入などの作業が必要
	// 変数　:= new(型)
	taro := new(MyData)
	fmt.Println(taro) // &{ []}
	taro.Name = "Taro"
	// make: 値を作成し、その初期化を行う
	// 配列・スライス、マップ、チャネルのみmakeを使える
	// 変数 = make(型, 個数)
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro) // &{Taro [0 0 0 0 0]}
}

// 構造体にメソッドを追加する
// func (割り当てる型の指定(レシーバ)) 関数名 (引数) 戻り値 {...処理の内容...}
func (md MyData) PrintData() {
	fmt.Println("*** MyData ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

func LessonStructMethod() {
	taro := MyData{
		"Hanako", []int{98, 76, 54, 32, 10},
	}
	taro.PrintData()
}

// 既存の型を拡張する
// intをintpという型でも使用できるようにした
type intp int

// intp型にメソッドを追加
func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n/2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// intp型にメソッドを追加
func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x %n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}

	ar = append(ar, x)
	return ar
}

func LessonExpansion() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n) // stringをintpに変換
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x *= 2
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}