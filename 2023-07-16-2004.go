package main

import "fmt"

func main() {
	// 1. 使用循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 使用循环计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 3. 使用循环打印出1到100之间的奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 4. 使用循环打印出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("")
	}

	// 5. 使用循环打印出斐波那契数列的前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}