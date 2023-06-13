package main

import "fmt"

func main() {
	// 练习1：打印1到10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1到10的和为：", sum)

	// 练习3：将一个字符串反转
	str := "hello world"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 练习4：使用递归计算n的阶乘
	n := 5
	fmt.Println(n, "的阶乘为：", factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}