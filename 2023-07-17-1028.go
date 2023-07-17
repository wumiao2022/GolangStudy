package main

import "fmt"

func main() {
	// 第一题：输出Hello World
	fmt.Println("Hello World")

	// 第二题：计算1到100的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 第三题：判断一个数是奇数还是偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 第四题：输出一个九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "*", j, "=", i*j, "\t")
		}
		fmt.Println()
	}
}