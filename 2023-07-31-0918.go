package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：计算两个浮点数的乘积
	num3 := 2.5
	num4 := 3.5
	product := num3 * num4
	fmt.Println("Product:", product)

	// 练习4：判断一个数是否为偶数
	num5 := 7
	if num5%2 == 0 {
		fmt.Println(num5, "is even")
	} else {
		fmt.Println(num5, "is odd")
	}

	// 练习5：求1到100的和
	sum2 := 0
	for i := 1; i <= 100; i++ {
		sum2 += i
	}
	fmt.Println("Sum of 1 to 100:", sum2)
}