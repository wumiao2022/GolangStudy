package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印
	sum := 1 + 2
	fmt.Println("1 + 2 =", sum)

	// 练习3：打印1到10之间的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：定义一个切片，并将切片中的值打印出来
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习5：交换两个变量的值并打印
	a, b := 10, 20
	fmt.Println("Before swap: a =", a, ", b =", b)
	a, b = b, a
	fmt.Println("After swap: a =", a, ", b =", b)
}