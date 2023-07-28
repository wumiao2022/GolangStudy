package main

import "fmt"

func main() {
	// 第一题：计算两个数的和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 第二题：将字符串拼接起来并输出结果
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	// 第三题：判断一个数是否为偶数，并输出结果
	num := 7
	if num%2 == 0 {
		fmt.Println("该数为偶数")
	} else {
		fmt.Println("该数为奇数")
	}

	// 第四题：循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 第五题：递归函数计算斐波那契数列的第n个数并输出结果
	n := 10
	fibonacci := fibonacci(n)
	fmt.Println(fibonacci)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}