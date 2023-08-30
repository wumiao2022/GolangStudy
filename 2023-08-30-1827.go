package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到10的所有整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：打印出1到100之间所有的奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习4：打印出1到100之间所有的偶数
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5：打印出100到1之间所有的整数（倒序）
	for i := 100; i >= 1; i-- {
		fmt.Println(i)
	}
}
