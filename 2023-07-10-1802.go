package main

import "fmt"

func main() {
	// 练习1：打印输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量，并打印输出其值
	var num int = 10
	fmt.Println(num)

	// 练习3：声明一个字符串变量，赋予你的名字，并打印输出
	var name string = "Alice"
	fmt.Println(name)

	// 练习4：计算两个整数的和，并打印输出
	var a int = 5
	var b int = 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习5：声明一个布尔变量，判断其是否为真，并打印输出
	var isTrue bool = true
	fmt.Println(isTrue)

	// 练习6：循环打印输出1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习7：使用if语句判断一个数是正数还是负数，并打印输出
	num = -5
	if num > 0 {
		fmt.Println("Positive")
	} else if num == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}

	// 练习8：编写一个函数，计算两个数的乘积，并返回结果
	result := multiply(4, 6)
	fmt.Println("Product:", result)
}

func multiply(a, b int) int {
	return a * b
}