package interview

import (
	"fmt"
	"reflect"
	"testing"
)

type MyEmptyInterface interface {
}

type NilFather interface {
	Show()
}

type NilSon struct {
}

func (item *NilSon) Show() {

}

// 这个地方：是否返回nil是interface的高阶知识点
func getFather() NilFather {
	//注：这里返回的是father的类型，而不是son的类型
	//区别是：这里返回的nilSon指针为nil
	//但是，该nilSon 指针会用于创建NilFather 接口，故整个方法返回的接口不为 nil
	var nilSon *NilSon
	return nilSon
}

func getSelf() *NilSon {
	//这种直接返回引用自己，所以返回的也是确定的nil
	var nilSon *NilSon
	return nilSon
}

// MyEmptyInterface是我们自己定义的空interface，无任何内部方法
func GetMyEmptyInterface() MyEmptyInterface {
	//虽然，nilSon 指针为nil
	//但是，该nilSon指针用于创建MyEmptyInterface接口，故，该接口不为nil
	var nilSon *NilSon
	return nilSon
}

// 直接返回空interface
func GetDefaultEmptyInterface() interface{} {
	//该nilSon指针用于创建GetDefaultEmptyInterface接口，故，该接口不为nil
	var nilSon *NilSon
	return nilSon
}

func TestNil(t *testing.T) {
	//这种直接声明但是没有初始化，直接声明的指针，是nil
	var a *NilSon
	fmt.Println(a == nil) //true

	//这种将子引用又赋值给父类，相当于重新构造父指针（父指针的类型不是nil）,故，其不为nil
	var b NilFather = a
	fmt.Println(b == nil) //false

	//这种是父类引用指向子对象,上面情况的变体，结果也是false
	father := getFather()
	fmt.Println(father == nil) //false

	father2 := getSelf()
	fmt.Println("getSelf:", father2 == nil) //true

	//如何判断nil接口的真实值是否为nil,使用反射直接获取其内部的真实数据，再判断是否为nil
	bIsNil := reflect.ValueOf(b).IsNil()
	fmt.Println(bIsNil) //true

	valueIsNil := reflect.ValueOf(father).IsNil()
	fmt.Println(valueIsNil) //true

	other := GetMyEmptyInterface()
	fmt.Println("GetMyEmptyInterface:", other == nil) //false

	other2 := GetDefaultEmptyInterface()
	fmt.Println("GetDefaultEmptyInterface:", other2 == nil) //false

}

func superJiekou2(i *interface{}) {
	//%v打印value, %T打印类型
	fmt.Printf("value %v, type=%T \n", i, i)
}

func TestXingInterface(t *testing.T) {
	var value *interface{}

	var str = "lisi"
	//step1: 将str转为 interface接口类型
	var x interface{} = str

	//step2: 然后，将interface转为指针 赋值给 *interface{}变量
	//注：不能直接将 str取地址赋值给 *interface{}变量，否则会报错
	//因为 *string 类型 != *interface{} 类型
	value = &x

	fmt.Println(*value)
}
