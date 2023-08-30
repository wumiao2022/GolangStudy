package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	// 3. 计算两个数的和并输出
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否是偶数
	num := 25
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 5. 输出1到100之间的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}