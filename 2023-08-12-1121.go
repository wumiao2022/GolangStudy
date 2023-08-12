package main

import "fmt"

func main() {
	// 练习1：打印输出Hello, Go!
	fmt.Println("Hello, Go!")

	// 练习2：输出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算并输出1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：交换两个变量的值，并输出交换后的结果
	a, b := 10, 20
	fmt.Println("Before swap:", a, b)
	a, b = b, a
	fmt.Println("After swap:", a, b)
}