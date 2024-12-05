package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestForrange(t *testing.T) {
	var list = make([]string, 0)

	list = append(list, "1")
	list = append(list, "2")
	list = append(list, "3")

	strs := make([]*string, 0)

	for _, item := range list {
		//item的地址是固定的还是不固定的取决于 go版本，更准确是 go.mod设置的版本
		fmt.Printf("%p, %v \n", &item, item)
		strs = append(strs, &item)
	}

	fmt.Println("----------------------")
	//验证最终结果
	for _, str := range strs {
		fmt.Printf("%p, %v \n", str, *str)
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
		//item++
	}

	fmt.Println("----------------------")
	for _, item := range funcList {
		//item是函数，执行之
		//item函数的执行逻辑，其实打印print的是上面for 循序item地址指向的内容
		//前面for循环结束后，item地址指向的内容也固定成了for循环的最后一个
		item()
	}
}

func TestForRangeVarAddress(t *testing.T) {
	//case1()
	case2()
}

func case1() {
	arr := []int{1, 2, 3}
	for _, val := range arr {
		go func() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
}

//输出结果：3 3 3g

func case2() {
	arr := [2]int{1, 2}
	arrnew := []*int{}
	for _, v := range arr {
		//场景1，使用for range循环变量地址构造其他数据
		//go <1.22 v只会创建同一个变量，故&v永远是相同地址
		fmt.Println(&v)
		arrnew = append(arrnew, &v)
	}

	for _, item := range arrnew {
		fmt.Printf("%p, %v \n", item, *item)
	}
}
