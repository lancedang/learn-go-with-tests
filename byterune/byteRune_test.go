package byterune

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	byte1 := 'A' //这种简写形式：是rune类型，默认是rune-占4个字节
	rune1 := 'B' //字符有2种形式：字节byte和rune类型

	fmt.Printf("byte size=%d \n", unsafe.Sizeof(byte1))
	fmt.Printf("rune size=%d \n", unsafe.Sizeof(rune1))

	var b2 byte = '1' //显示声明-占1个字节
	var r2 rune = '2' //

	fmt.Printf("byte size=%d \n", unsafe.Sizeof(b2))
	fmt.Printf("rune size=%d \n", unsafe.Sizeof(r2))
}

func TestVar(t *testing.T) {
	a := "xx"
	b := a //2份数据，更改彼此不影响

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
