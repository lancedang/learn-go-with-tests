package interview

import (
	"fmt"
	"testing"
)

func defer1() {
	fmt.Println("1")
}

func defer2() {
	fmt.Println("2")
}

func TestDefer(t *testing.T) {
	//定义顺序决定压栈顺序
	//先定义后执行，先进后出
	defer defer1()
	defer defer2()
}

func defer3() int {
	fmt.Println("defer 3")
	return 3
}

func returnFunc4() int {
	fmt.Println("return 4")
	return 4
}

func get() int {
	//defer关键字修饰的执行顺序上：总是最后执行（晚于return的函数）
	//但是，最终的结果，还是return 后的函数，即defer的函数的返回值无效
	defer defer3()
	return returnFunc4()
}

func TestDeferAndReturn(t *testing.T) {
	i := get()
	fmt.Println("result=", i)
}

func DeferWithReturn() (result int) {

	defer func() {
		fmt.Println("defer非函数里直接更改结果值result=100")
		result = 100
	}()

	return 10
}

func TestPrintReturn(t *testing.T) {
	result := DeferWithReturn()
	fmt.Println("外部调用方法结果=", result)
}

func deferWithReturn1() int {
	fmt.Println("defer函数的return返回1，但是无效")
	return 1
}

func deferNoReturn() {
	fmt.Println("defer函数里没有return返回值")
}

func returnMethod() int {
	fmt.Println("return函数的return返回10，有效")

	return 10
}

func deferWithReturnDemo() int {
	defer deferWithReturn1()
	return returnMethod()
}

func TestDeferWithReturn(t *testing.T) {
	result := deferWithReturnDemo()
	fmt.Println("最终结果：", result)
}

// 测试单个返回值 string vs （xx string）区别， defer里 xx== 10这种写法
func deferWithParam10(param int) {
	fmt.Println("int类型设置：defer函数里重新设置参数值,param=10")
	param = 10
}

func returnDemoWithParam(param int) (result int) {
	defer deferWithParam10(param)

	fmt.Println("int类型设置：defer函数外设置param=20")
	param = 20

	return param
}

func TestReturnWithParam(t *testing.T) {
	result := returnDemoWithParam(1)
	fmt.Println("int类型设置：result=", result)
}

func deferWithPointerParam10(param map[string]int) {
	fmt.Println("map类型设置：defer函数里重新设置参数值,param=10")
	param["data"] = 10
}

func returnDemoWithPointerParam(param map[string]int) (result map[string]int) {
	defer deferWithPointerParam10(param)

	fmt.Println("map类型设置：defer函数外设置param=20")
	param["data"] = 20

	return param
}

func TestReturnWithPointerParam(t *testing.T) {
	mapper := make(map[string]int)
	result := returnDemoWithPointerParam(mapper)
	fmt.Println("map类型设置：result=", result["data"])
}

func changeInt(x int) {
	fmt.Println("change方法重设int值x=10")
	x = 10
}

func TestChangeIntByParam(t *testing.T) {
	x := 1
	//change方法里改变int参数x,不影响返回结果
	changeInt(x)

	fmt.Println("最终结果=", x)
}

func changeMap(mapp map[string]int) {
	fmt.Println("change方法重设map某个key=10")
	mapp["data"] = 10
}

func TestChangeMapByParam(t *testing.T) {
	mapp := make(map[string]int)
	mapp["data"] = 20
	//change方法里改变map参数x,会影响返回结果
	changeMap(mapp)
	fmt.Println("最终结果=", mapp["data"])
}
