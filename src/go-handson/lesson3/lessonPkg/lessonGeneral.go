package lessonPkg

import "fmt"

// 空のinterface
type General interface{}

// Generalの値を保持するinterface
type GData interface {
	Set(nm string, g General) GData
	Print()
}

// Generalの値を使ったメソッドを実装した構造体
// type GDataImpl struct {
// 	Name string
// 	Data General
// }

// func (gd * GDataImpl) Set(nm string, g General) {
// 	gd.Name = nm
// 	gd.Data = g
// }

// func (gd *GDataImpl) Print() {
// 	fmt.Printf("<<%s>>", gd.Name)
// 	fmt.Println(gd.Data)
// }

// func LessonGeneral() {
// 	var data = []GDataImpl{}
// 	// Generalはどんな値でも補完できる
// 	data = append(data,GDataImpl{"Hanako", 123})
// 	data = append(data,GDataImpl{"Hanako", "hello!"})
// 	data = append(data,GDataImpl{"Hanako", []int{123, 456, 789}})
// 	for _, ob := range data {
// 		ob.Print()
// 	}
// }


/*
 * 型アサーション
 * 変数 := 値 .(型)
*/

type NData struct {
	Name string
	Data int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	nd.Data = g.(int) // 型アサーションを実施 (空のinterfaceに型を指定？)
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	sd.Data = g.(string) // 型アサーションを実施 (空のinterfaceに型を指定？)
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func LessonAssertion() {
	var data = []GData{}
	data = append(data,new(NData).Set("Taro", 123))
	data = append(data,new(SData).Set("Jiro", "hello!"))
	data = append(data,new(NData).Set("Hanako", 789))
	data = append(data,new(SData).Set("Sachiko", "hello?"))
	for _, ob := range data {
		ob.Print()
	}
}