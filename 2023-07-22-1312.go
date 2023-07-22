package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印
	result := 1 + 2
	fmt.Println(result)

	// 练习3：检查一个数字是否是偶数
	num := 4
	isEven := num%2 == 0
	fmt.Println(isEven)

	// 练习4：将一个整数转换为字符串并打印
	numStr := fmt.Sprintf("%d", 123)
	fmt.Println(numStr)

	// 练习5：使用循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}