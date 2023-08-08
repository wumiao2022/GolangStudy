package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num3 := 7
	if num3%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：使用循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：计算并打印1到100的所有偶数的和
	sumEven := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sumEven += i
		}
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sumEven)
}
