package copy

import (
	"fmt"
	"testing"
)

func TestMapValue(t *testing.T) {
	a1 := A{
		name: "lisi",
	}

	//一、构造普通对象类型(非指针)类型的value的map
	aMap := map[string]A{}
	aMap["data1"] = a1 //附一个非指针value

	fmt.Println("old=", aMap)
	//直接访问map的value是不ok的->map之value不可寻址性
	//故，此行连编译都通不过
	//aMap["data1"].name = "张三"

	//必须拿个变量中转下，才能访问,且中转是值拷贝（而非引用拷贝），修改彼此不影响
	tempA1 := aMap["data1"]
	tempA1.name = "张三"
	fmt.Println("new=", aMap)
	fmt.Println("-----------")

	//二、构造指针类型的value的map
	aMap2 := map[string]*A{}
	aMap2["data1"] = &a1 //赋值一个指针类型

	fmt.Println("old=", aMap2["data1"])
	//如下，可以访问，编译可通过，且改变彼此影响，改动的是同一份数据
	aMap2["data1"].name = "这次改变了吧"
	fmt.Println("new=", aMap2["data1"])
}
