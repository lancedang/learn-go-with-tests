package copy

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	a := "xx"
	b := a //复制一份数据，2份数据，更改彼此不影响，这个:=是复制的操作

	b = "yy"

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("-----------")

	a2 := "xx"
	b2 := &a2 //取数据的地址

	*b2 = "yy" //给数据地址处 赋值， *b2代表b2地址那的数据用"yy"来填充，更改其一，彼此影响

	fmt.Println(a2)
	fmt.Println(*b2)
}
