package main

import "fmt"

func main() {
	// 练习1: 使用for循环打印出1到10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 使用for循环打印出从2到10的偶数
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习3: 使用for循环计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4: 使用for循环打印出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}