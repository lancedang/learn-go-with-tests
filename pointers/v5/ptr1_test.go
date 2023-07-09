package v5

import (
	"fmt"
	"testing"
)

func TestPtr(t *testing.T) {
	age := "18" //变量-存储数据

	agePtr := &age //&-标识指针变量-取地址-存储数据的地址

	ageFromAddr := *agePtr //*-这个不是指针-取某个地址(&修饰的变量)处的数据-又叫解指针

	fmt.Printf("ageFromAddr=%s \n", ageFromAddr)

	fmt.Printf("age=%s \n", age)

	fmt.Printf("age ptr=%p \n", agePtr)
	fmt.Println("age ptr=", agePtr)
}
