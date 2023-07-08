package main

import "fmt"

func main() {
	// 1. 声明一个整型变量x，并赋值为10
	var x int = 10

	// 2. 声明一个字符串变量name，并赋值为"Go"
	var name string = "Go"

	// 3. 声明一个浮点型变量pi，并赋值为3.14
	var pi float64 = 3.14

	// 4. 打印变量x的值
	fmt.Println(x)

	// 5. 打印变量name的值
	fmt.Println(name)

	// 6. 打印变量pi的值
	fmt.Println(pi)

	// 7. 将变量x的值加上5，并将结果存储回变量x
	x = x + 5

	// 8. 将变量name的值与" is awesome!"拼接起来，并将结果存储回变量name
	name = name + " is awesome!"

	// 9. 将变量pi的值乘以2，并将结果存储回变量pi
	pi = pi * 2

	// 10. 打印变量x的值
	fmt.Println(x)

	// 11. 打印变量name的值
	fmt.Println(name)

	// 12. 打印变量pi的值
	fmt.Println(pi)
}
