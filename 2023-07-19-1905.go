package main

import "fmt"

func main() {
	// 练习1：打印 1 到 10 的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印从 10 到 1 的所有整数
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习3：打印 1 到 10 的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：计算并打印 1 到 10 的所有整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5：计算并打印 1 到 10 的所有整数的乘积
	product := 1
	for i := 1; i <= 10; i++ {
		product *= i
	}
	fmt.Println(product)
}