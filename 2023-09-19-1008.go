package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求和函数
	sum := add(3, 4)
	fmt.Println("3 + 4 =", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 5
	if isEven(num) {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习4：计算给定数的阶乘
	factorial := calcFactorial(5)
	fmt.Println("5的阶乘是：", factorial)
}

// 练习2：求和函数
func add(a, b int) int {
	return a + b
}

// 练习3：判断一个数是偶数还是奇数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 练习4：计算给定数的阶乘
func calcFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcFactorial(n-1)
}
