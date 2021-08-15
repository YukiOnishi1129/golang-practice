package lessonPkg

import "fmt"

/*
  type 名前 interface {
	  メソッドA
	  メソッドB
	  ...必要なだけ用意.....
  }
*/

// Data is interface
// InitialとPrintDataInterfaceを持つ構造体は、
// 自動的にDataインターフェイスを実装しているとみなされる
type Data interface {
	Initial(name string, data []int)
	PrintDataInterface()
}


type MyDataInter struct {
	Name string
	Data []int
}

func (md *MyDataInter) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *MyDataInter) PrintDataInterface() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func LessonInterface() {
	// MyDataInterの値をData型の変数に代入
	// newは代入する変数の方に合わせて値が扱われる
	var ob Data = new(MyDataInter)
	// エラーになる (MyDataInterの値をData型に入れようとするため)
	// var ob Data = MyDataInter{}
	// エラーになる (interfaceに直接値を入れることはできない)
	// var ob Data = Data{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintDataInterface()
}


func (md * MyDataInter) Check() {
	fmt.Printf("Check! [%s]", md.Name)
}

func LessonInterface2() {
	var ob MyDataInter = MyDataInter{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	// 問題なく呼び出せる
	ob.Check()
}

func LessonInterface3() {
	var ob Data = new (MyDataInter)
	ob.Initial("Sachiko", []int{55, 66, 77})
	// エラーになる
	// ob.Check undefined (type Data has no field or method Check)
	// newしてるのはMyDataInterで値(Check)もあるはずだが、
	// 大事なのは「型」なので、エラーになる
	// Data型の値として振る舞うので、Checkは定義されていないのでエラー
	// インターフェース型として変数を扱う場合は、そのインターフェースに用意された機能しか使えない
	// ob.Check()
}