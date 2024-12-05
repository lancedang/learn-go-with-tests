package interview

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

func TestModifyMapByZhiyinyong(t *testing.T) {

	//值类型value
	m := make(map[string]Student)
	student := Student{Name: "is"}

	m["t1"] = student

	//value是非指针类型，以下编译报错，不能直接赋值
	//只有pointer类型的value才可以直接赋值
	//m["t1"].Name = "xx"

	fmt.Printf("初始student: %v\n", m["t1"])

	//修改t1的key对于的value
	//相当于复制了一份value对象，有点Java的写时赋值的概念
	//更新操作是在复制出来的副本上操作的，故除非写回map,否则更新只存在副本上
	temStudent := m["t1"]
	temStudent.Name = "zhangsan"

	fmt.Printf("重新设置之前: %v\n", m["t1"])

	//必须重新设置回去，更新才生效
	m["t1"] = temStudent

	fmt.Printf("重新设置之后: %v\n", m["t1"])
}

func TestModifyMapByPointerValue(t *testing.T) {
	m := make(map[string]*Student)
	student := &Student{Name: "lisi"}

	m["t1"] = student
	fmt.Printf("初始student: %v\n", m["t1"])

	student.Name = "zhangsan"
	fmt.Printf("通过value指针 重新设置之后: %v\n", m["t1"])

	//因为map的value为pointer类型，故以下操作可以直接赋值
	m["t1"].Name = "wangwu"
	fmt.Printf("通过map key直接设置value属性: %v\n", m["t1"])

	//因为map的value为pointer类型，故取出的value后直接操作即可，不需重新设置回去
	tempStudent := m["t1"]
	tempStudent.Name = "zhaoliu"
	fmt.Printf("通过map key获取value对象，重新设置之后: %v\n", m["t1"])

}

func TestCreateMapByArray(t *testing.T) {
	students := make([]Student, 0)

	students = append(students, Student{Name: "1"})
	students = append(students, Student{Name: "2"})
	students = append(students, Student{Name: "3"})

	for _, item := range students {
		fmt.Printf("student: add=%p %v\n", &item, item)
	}

	strs := make([]string, 0)

	strs = append(strs, "1")
	strs = append(strs, "2")
	strs = append(strs, "3")

	for _, item := range strs {
		fmt.Printf("string: add=%p %v\n", &item, item)
	}
}
