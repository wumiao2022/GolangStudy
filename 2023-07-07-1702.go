package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数相加的结果，并打印出来
	num1, num2 := 10, 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：判断一个整数是否为偶数，并打印出结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：使用循环打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用循环计算1到10的累加和，并打印出结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The sum is:", sum)
}