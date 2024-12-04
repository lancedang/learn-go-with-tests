package interview

import (
	"fmt"
	"testing"
)

func TestForrange(t *testing.T) {
	var list = make([]string, 0)

	list = append(list, "1")
	list = append(list, "2")
	list = append(list, "3")

	for i, item := range list {
		fmt.Println(i, &i, &item, ",", item)
	}
}

func TestVarFlee(t *testing.T) {
	var funcList []func()
	intList := []int{1, 2, 3}

	for _, item := range intList {
		fmt.Println("item1=", item)
		//innerItem := item
		//在整个for循环中，item地址是相同的固定的某一个
		funcList = append(funcList, func() {
			//在这个函数里其实，是将item的地址传递出去了，传递到for循环的外部了
			//即，item的生命周期发生了变化（栈->堆）也就是所谓的变量逃逸
			fmt.Println("item2=", item)
		})

		//若修改item,其实影响逃逸后的真实值的
		item++
	}

	//
	for _, item := range funcList {
		//item是函数，执行之
		//item函数的执行逻辑，其实打印print的是上面for 循序item地址指向的内容
		//前面for循环结束后，item地址指向的内容也固定成了for循环的最后一个
		item()
	}

}
