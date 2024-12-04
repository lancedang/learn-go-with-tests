package interview

import (
	"fmt"
	"testing"
)

type Father struct {
}

type Mather struct {
}

// 这种方法对应java的非static普通方法，必须使用类的对象进行调用
func (father *Father) fatherEat() {
	fmt.Println("father fatherEat")
}

func (mather *Mather) matherSong() {
	fmt.Println("mather matherSong")
}

// 这类方法类似java的static方法，不隶属于某个类，直接调用即可
func show() {

}

type Son struct {
	Father
	Mather
}

func (son *Son) sonWrite() {
	fmt.Println("son sonWrite")
}

func TestExtend(t *testing.T) {
	var son = Son{
		Father: Father{}, //引入第一个其他类
		Mather: Mather{}, //引入第二关其他类，从而实现多继承
	}

	//调用继承的第一个方法
	son.fatherEat()
	//调用继承的第二个方法
	son.matherSong()
	//调用自己的方法
	son.sonWrite()

}
