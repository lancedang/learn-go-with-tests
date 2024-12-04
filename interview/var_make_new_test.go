package interview

import (
	"fmt"
	"testing"
)

func TestCreateVariable(t *testing.T) {
	var x1 = "1"
	var x2 = new(string)

	//
	fmt.Println("var xx=\"\" 的直接就是值，不是指针：", x1)

	fmt.Println("new()这是个地址：", x2)
	fmt.Println("new()使用*获取地址指向的默认值：", *x2)

	*x2 = "啥"
	fmt.Println("给引用变量赋值：", *x2)

	var slice1 = make([]int, 0)
	fmt.Println(slice1)
	fmt.Println(&slice1)

}
