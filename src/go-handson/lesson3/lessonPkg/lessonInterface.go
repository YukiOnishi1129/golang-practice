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
	var ob MyDataInter = MyDataInter{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintDataInterface()
}