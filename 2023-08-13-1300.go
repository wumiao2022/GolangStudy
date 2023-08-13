package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, Go!
	fmt.Println("Hello, Go!")

	// 练习2：计算 7 除以 3 的余数并打印结果
	fmt.Println(7 % 3)

	// 练习3：声明一个字符串变量 name，并将其赋值为你的名字，并打印出来
	name := "John"
	fmt.Println(name)

	// 练习4：声明一个整数变量 age，并将其赋值为你的年龄，并打印出来
	age := 30
	fmt.Println(age)

	// 练习5：声明一个浮点数变量 pi，并将其赋值为 3.1415926，并打印出来
	pi := 3.1415926
	fmt.Println(pi)

	// 练习6：声明一个布尔型变量 isFemale，并将其赋值为 true，并打印出来
	isFemale := true
	fmt.Println(isFemale)

	// 练习7：声明一个切片变量 numbers，包含元素 1、2、3，并打印出来
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// 练习8：声明一个字典变量 person，包含键值对 name: "John"、age: 30，并打印出来
	person := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	fmt.Println(person)
}