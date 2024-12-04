package interview

import (
	"fmt"
	"testing"
)

// 定义接口
type SaySth interface {
	Say()
}

// 定义A类
type A struct {
	Name string
}

// A类实现上面接口方法
func (a *A) Say() {
	fmt.Printf("实现类执行：a say %v \n", a.Name)
}

type B struct {
	Name string
}

// B类也实现上面接口方法
func (b *B) Say() {
	fmt.Printf("实现类执行：b say %v \n", b.Name)
}

// 定义一个多态方法,方法参数是接口类型，注：接口不能使用引用*修饰
func RealSay(saySth SaySth) {
	//调用具体类实现多态
	fmt.Printf("接口调用中转：saySth地址=%p \n", &saySth)
	saySth.Say()
}

type C struct {
	Name string
}

// 看到没，这里的实现类c,不使用引用*号修饰也行，区分上面的 (a *A)
func (c C) Say() {
	fmt.Printf("实现类执行：c say %v \n", c.Name)
}

func TestDuotai(t *testing.T) {
	var a = A{Name: "lisi"}
	var b = B{Name: "7"}

	fmt.Printf("调用前，构造变量：a地址=%p \n", &a)
	fmt.Printf("调用前，构造变量：b地址=%p \n", &b)
	RealSay(&a)
	RealSay(&b)

	var c = C{Name: "shanghai"}
	//这个就不用&取地址（引用）了
	fmt.Printf("调用前，构造变量：c地址=%p \n", &c)
	RealSay(c)

	//这的写法严格按照 父引用指向子对象地址的写法，否则会报错
	//必须要有&A{} &修饰
	var aPlus SaySth = &A{Name: "a 北京"}
	aPlus.Say()

	var cPlus SaySth = &C{Name: "c 上海"}
	var cPlus2 SaySth = C{Name: "c 上海"}
	cPlus.Say()
	cPlus2.Say()
}
