package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：求两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 9
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算1到100之间的所有偶数和
	sumEven := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sumEven += i
		}
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sumEven)
}