package interview

import (
	"fmt"
	"testing"
)

func TestReturn(t *testing.T) {

}

func r1() (x int) {
	//可以不设置x,直接返回某个值10，这个值自动赋予x变量
	return 10
}

func r2() (x int) {
	x = 5
	return x
}

func r3() (x int) {
	y := x + 1
	return y
}

func rPlus1() int {
	return 10
}

func rPlus2() int {
	x := 1
	defer func() {
		//这种defer写法：在方法体返回标志非( x int) ，即只有(int)这种情况下，对x的修改不起作用
		x = 10
	}()

	x = 2
	return x
}

func rPlus2Plus() (x int) {

	defer func() {
		//这种defer里面直接修改x生效，会最终影响 函数返回结果
		x = 10
	}()

	x = 2
	return x
}

func TestReturnNoName(t *testing.T) {
	plus2 := rPlus2()
	fmt.Println("未指定返回时的接收参数，只有返回类型：", plus2)

	plus3 := rPlus2Plus()
	fmt.Println("同时指定返回时的接收参数和类型，：", plus3)
}
