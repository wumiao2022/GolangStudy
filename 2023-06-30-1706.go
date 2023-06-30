package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 练习2：计算两个整数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 练习3：打印1到10的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
	
	// 练习4：打印1到10的偶数
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
	
	// 练习5：打印1到100的倍数
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
	}
}