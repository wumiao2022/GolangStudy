package main

import "fmt"

func main() {
	// 练习1：打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算并打印1到5的和
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：打印1到10的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：计算并打印5的阶乘
	factorial := 1
	for i := 1; i <= 5; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	// 练习5：打印一个由*组成的正方形
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}