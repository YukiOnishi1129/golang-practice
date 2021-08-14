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