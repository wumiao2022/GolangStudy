package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个名为name的字符串变量，并将其值设置为你的名字，然后打印出来
	name := "John"
	fmt.Println("My name is", name)

	// 练习3：定义一个整数变量x，将其值设置为10，然后打印出来
	x := 10
	fmt.Println("x is", x)

	// 练习4：定义一个函数add，接受两个整数参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 练习5：调用add函数，并打印返回结果
	result := add(5, 3)
	fmt.Println("The sum of 5 and 3 is", result)

	// 练习6：循环打印数字1到10的平方值
	for i := 1; i <= 10; i++ {
		fmt.Println("The square of", i, "is", i*i)
	}
}