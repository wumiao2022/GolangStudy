package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数并输出结果
	num := 17
	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}

	// 练习4：打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用for循环计算并输出1到10的所有整数的平方
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Squared:", i*i)
	}
}