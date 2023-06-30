package main

import "fmt"

func main() {
	// 打印 Hello World
	fmt.Println("Hello World!")

	// 打印 1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 计算两个数的和并打印结果
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 打印一个由数字组成的倒三角形
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}