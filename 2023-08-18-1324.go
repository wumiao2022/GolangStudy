package main

import "fmt"

func main() {
	// 练习1：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}