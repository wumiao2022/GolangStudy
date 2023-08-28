package main

import "fmt"

func main() {
	// 练习1：打印"Hello, world!"
	fmt.Println("Hello, world!")

	// 练习2：计算两个整数相加的结果并打印
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 17
	isEven := num%2 == 0
	fmt.Println("Is", num, "even?", isEven)

	// 练习4：打印1到10的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习5：计算1到100的和并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100:", sum)
}