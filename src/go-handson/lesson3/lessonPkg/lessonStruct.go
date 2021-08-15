package lessonPkg

import "fmt"

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