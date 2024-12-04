package interview

import (
	"fmt"
	"testing"
)

type ZhuanHuan struct {
	Name string
}

// 情况1： 定义值类型的接收者(这里不是方法参数，是方法所隶属的类struct)
func (x ZhuanHuan) checkInReceiver() {
	fmt.Println("方法隶属的类型-结果:", x.Name)
}

// 情况2：值类型放在方法参数里
func checkInMethodParam(x ZhuanHuan) {
	fmt.Println("方法参数-结果:", x.Name)
}

func TestReceive(t *testing.T) {
	var x1 = ZhuanHuan{Name: "x1"}
	//这个地方定义x2是指针类型
	var x2 = &ZhuanHuan{Name: "x2"}

	x1.checkInReceiver()
	//x2是指针类型，为何还能像上面值类型一样，直接调用这个函数？
	//这种没有编译报错
	x2.checkInReceiver()

	checkInMethodParam(x1)
	//这种是方法参数接收的值类型，则必须传递 值，不能传递引用类型 ，否则编译报错
	//checkInMethodParam(x2)

}

// 问题：这个CommonFather接口可以接收任何类型的返回值
type CommonFather interface {
	doSth()
}

type CommonSon struct {
}

func (x *CommonSon) doSth() {

}

func TestFatherPoint2Son(t *testing.T) {
	var father CommonFather = &CommonSon{}
	father.doSth()
}

// i可以接收任何类型的实际参数
func superJiekou(i interface{}) {
	//%v打印value, %T打印类型
	fmt.Printf("value %v, type=%T \n", i, i)
}

func TestSuperJiekou(t *testing.T) {
	superJiekou("张三")
	superJiekou(100)
	superJiekou(1.23)
	superJiekou(CommonSon{})

}

func TestDuanYanJiekouType(t *testing.T) {
	var i interface{} = 2

	//写法1-直接判断
	if v, ok := i.(int); ok {
		fmt.Printf("%v is int \n", v)
	} else if v, ok := i.(string); ok {
		fmt.Printf("%v is string \n", v)
	}

	//写法2-switch type判断，这种必须在Switch语法里
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int \n", v)
	case string:
		fmt.Printf("%v is int \n", v)
	default:
		fmt.Printf("unknow type \n")
	}
}
