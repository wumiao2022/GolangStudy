package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 使用for循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3: 计算两个数的和并打印结果
	a := 10
	b := 5
	sum := a + b
	fmt.Println(sum)

	// 练习4: 判断一个数是否为奇数并打印结果
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5: 使用切片存储一组字符串并打印出来
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva"}
	for _, name := range names {
		fmt.Println(name)
	}

	// 练习6: 创建一个函数，接收两个整数参数并返回它们的乘积
	fmt.Println(multiply(4, 5))
}

func multiply(a, b int) int {
	return a * b
}