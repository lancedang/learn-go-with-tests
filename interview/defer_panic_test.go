package interview

import (
	"fmt"
	"testing"
)

func TestDeferWithPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("最终捕获的是代码执行顺序的最后那个panic：", err)
		}
	}()

	defer func() {
		//这个后发生
		fmt.Println("defer操作1")
		panic("panic发生在defer-1里")
	}()

	defer func() {
		//这个后发生
		fmt.Println("defer操作2")
		panic("panic发生在defer-2里")
	}()

	//这个先发生，
	panic("正常逻辑的panic")
}
