package lessonPkg

import (
	"fmt"
	"reflect"
)

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
	Data []int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	// 型をチェック
	// reflect.TypeOf(値)
	// reflect.TypeOf(値).Kind()で値の型を取り出す

	// int型の配列かチェック
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf((0))) {
		nd.Data = g.([]int) // 型アサーションを実施 (空のinterfaceに型を指定？)
	}
	
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data []string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	// 型をチェック
	// reflect.TypeOf(値)
	// reflect.TypeOf(値).Kind()で値の型を取り出す
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = g.([]string) // 型アサーションを実施 (空のinterfaceに型を指定？)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

// func LessonAssertion() {
// 	var data = []GData{}
// 	data = append(data,new(NData).Set("Taro", 123))
// 	data = append(data,new(SData).Set("Jiro", "hello!"))
// 	data = append(data,new(NData).Set("Hanako", 789))
// 	data = append(data,new(SData).Set("Sachiko", "hello?"))
// 	for _, ob := range data {
// 		ob.Print()
// 	}
// }

// func LessonReflect() {
// 	var data = []GData{}
// 	data = append(data,new(NData).Set("Taro", 123))
// 	data = append(data,new(SData).Set("Jiro", "hello!"))
// 	data = append(data,new(NData).Set("Hanako", "89766"))
// 	data = append(data,new(SData).Set("Sachiko", []string{"happy?"}))
// 	for _, ob := range data {
// 		ob.Print()
// 	}
// }


/*
  配列をreflect.TypeOfでチェックする方法
  配列のTypeを得る: reflect.ArrayOf(<Type>)
  スライスのTypeを得る: reflect.SliceOf(<Type>)
  マップのTypeを得る: reflect.MapOf(<Type>)
*/

func LessonArrayTypeCheck() {
	var data = []GData{}
	data = append(data,new(NData).Set("Taro", []int{1, 2, 3}))
	data = append(data,new(SData).Set("Jiro", []string{"Hello", "bye"}))
	data = append(data,new(NData).Set("Hanako", 876))
	data = append(data,new(SData).Set("Sachiko", "happy?"))
	for _, ob := range data {
		ob.Print()
	}
}