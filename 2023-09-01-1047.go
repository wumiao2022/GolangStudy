package main

import "fmt"

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 10:", sum)

	// 练习3：打印10到1的倒序数
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习4：打印1到100之间的奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5：打印1到100之间的偶数
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}
}