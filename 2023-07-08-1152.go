package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	sum := add(10, 20)
	fmt.Println("The sum is", sum)

	// 练习3：判断一个数是正数、负数还是零
	number := -5
	if number > 0 {
		fmt.Println(number, "is a positive number")
	} else if number < 0 {
		fmt.Println(number, "is a negative number")
	} else {
		fmt.Println("The number is zero")
	}

	// 练习4：使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 函数：计算两个整数的和
func add(a, b int) int {
	return a + b
}