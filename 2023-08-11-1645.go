package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 2. 计算两个整数的和并输出
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 3. 计算两个浮点数的积并输出
	num3 := 3.14
	num4 := 2.5
	product := num3 * num4
	fmt.Println("Product:", product)
	
	// 4. 判断一个数是否为偶数并输出结果
	num5 := 15
	if num5%2 == 0 {
		fmt.Println(num5, "is even")
	} else {
		fmt.Println(num5, "is odd")
	}
	
	// 5. 计算1到100的和并输出
	sum2 := 0
	for i := 1; i <= 100; i++ {
		sum2 += i
	}
	fmt.Println("Sum of 1 to 100:", sum2)
}