package main

import "fmt"

func main() {
	// 练习1：输出Hello World!
	fmt.Println("Hello World!")

	// 练习2：定义并输出两个整数变量的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：定义并输出一个字符串变量
	message := "Hello Golang!"
	fmt.Println(message)

	// 练习4：计算并输出10的阶乘
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Println("10! =", factorial)

	// 练习5：定义一个切片并输出其中的元素
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// 练习6：使用if-else语句输出一个数字的奇偶性
	number := 12
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习7：使用for循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}