package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：求两个整数之和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 12
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：打印1到100的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5：计算1到100的累加和
	sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100:", sum)
}