package copy

import (
	"fmt"
	"testing"
)

type A struct {
	name string
}

func TestArrayAndSliceCopy(t *testing.T) {
	arr := [2]int{1, 2}

	arr2 := arr //复这个在go中是复制一份数据的意思，拷贝的表现形式

	arr2[0] = 10 //修改第二份数据，不影响第一份数据-数组是值复制

	fmt.Println(arr)  //[1 2]
	fmt.Println(arr2) //[10 2]

	fmt.Println("-------------")

	slice := []int{1, 2}
	slice2 := slice //slice的复制，复制一份slice,指针的复制，修改彼此影响，拷贝的表现形式

	slice2[0] = 1000 //修改slice2，会影响slice1

	fmt.Println(slice)  //[1000 2] //也跟着变成最新的1000了
	fmt.Println(slice2) //[1000 2]

	aArray := [2]A{{
		name: "111",
	}, {
		name: "222",
	}}

	fmt.Println("----------")

	aArray2 := aArray

	aArray2[1].name = "xxxx"

	fmt.Println(aArray)
	fmt.Println(aArray2)

	fmt.Println("----------")

	aSlice := []A{{
		name: "111",
	}, {
		name: "222",
	}}

	aSlice2 := aSlice

	aSlice2[1].name = "yyyy"

	fmt.Println(aSlice)
	fmt.Println(aSlice2)

}

func TestPointer(t *testing.T) {
	a := A{name: "张三"}

	//定义指针类型
	var pointer *A = &a

	Show(pointer)
}

// Show方法接收的是指针类型的参数
func Show(x *A) {
	//step-1）获取指针表示的地址处的数据
	var item = *x

	//step-2）获取数据的字段
	fmt.Println(item.name)

	//上面2步，可以跳过获取 *x的过程，直接使用x来操作
	//原理是 .号帮我解析成上面的步骤1
	fmt.Println(x.name)
}
