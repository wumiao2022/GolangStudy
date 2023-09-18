package main

import "fmt"

func main() {
	// 1. 输出Hello World
	fmt.Println("Hello World")

	// 2. 打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 5. 打印斐波那契数列的前n个数
	n := 10
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}