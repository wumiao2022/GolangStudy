package main

import "fmt"

func main() {
	// 练习1: 声明一个整型变量x并初始化为10，打印出x的值
	x := 10
	fmt.Println(x)

	// 练习2: 声明一个字符串变量name并初始化为"John"，打印出name的值
	name := "John"
	fmt.Println(name)

	// 练习3: 声明一个布尔类型变量isTrue并初始化为true，打印出isTrue的值
	isTrue := true
	fmt.Println(isTrue)

	// 练习4: 声明一个浮点型变量pi并初始化为3.14，打印出pi的值
	pi := 3.14
	fmt.Println(pi)

	// 练习5: 声明一个整型变量a并初始化为5，声明一个整型变量b并初始化为3，计算a除以b的结果并打印出来
	a := 5
	b := 3
	result := a / b
	fmt.Println(result)

	// 练习6: 声明一个字符串变量message并初始化为"Hello, world!"，将其转换为大写并打印出来
	message := "Hello, world!"
	uppercaseMessage := strings.ToUpper(message)
	fmt.Println(uppercaseMessage)
}
