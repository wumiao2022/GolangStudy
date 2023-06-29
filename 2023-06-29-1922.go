package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两数之和
	fmt.Println(add(3, 5))

	// 练习3：判断一个数是奇数还是偶数
	checkNumber(7)

	// 练习4：循环打印数字
	printNumbers(10)
}

func add(a, b int) int {
	return a + b
}

func checkNumber(num int) {
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
