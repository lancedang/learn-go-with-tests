package interview

import (
	"fmt"
	"testing"
)

func open(first int, second int) int {
	fmt.Println(first)
	return first
}

func TestDeferWithSonFunction(t *testing.T) {
	//逻辑很简单：defer的所有函数从上到下依次进栈
	//且每个defer函数进栈的时候，其方法参数必须先计算出来才能进栈
	//故：
	//open(3,0)先计算结果->3
	//open(1,3)压栈
	//open(4,0) 也先计算结果->4
	//open(2,4)压栈

	//之后：两个压栈的开始执行 2->1
	defer open(1, open(3, 0))
	defer open(2, open(4, 0))
}
