package main

import "fmt"

func main() {
	// 练习1：打印1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印10~1
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习3：打印0~20中的偶数
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习5：计算并打印1~100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}