package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个整数相加
	num1 := 10
	num2 := 5
	result := num1 + num2
	fmt.Println("结果：", result)

	// 练习4：判断一个数是否为奇数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习5：使用函数计算阶乘
	num := 5
	factorial := calcFactorial(num)
	fmt.Println(num, "的阶乘为", factorial)
}

func calcFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcFactorial(n-1)
}
