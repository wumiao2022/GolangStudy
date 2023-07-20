package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	fmt.Println(2 + 3)

	// 练习3：判断一个数是否为偶数
	num := 4
	if num%2 == 0 {
		fmt.Println("这是一个偶数")
	} else {
		fmt.Println("这是一个奇数")
	}

	// 练习4：打印1到10的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习5：使用函数实现阶乘计算
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}