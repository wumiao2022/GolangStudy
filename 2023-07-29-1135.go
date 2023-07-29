package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印
	sum := 1 + 2
	fmt.Println("1 + 2 =", sum)

	// 练习3：将5赋值给变量x，然后打印x的值
	x := 5
	fmt.Println("x =", x)

	// 练习4：声明一个变量y并赋值为10，然后将y的值加上x，并将结果打印
	y := 10
	result := x + y
	fmt.Println("x + y =", result)

	// 练习5：声明一个常量pi并赋值为3.14159，然后打印pi的值
	const pi = 3.14159
	fmt.Println("pi =", pi)
}