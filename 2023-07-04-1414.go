package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：将两个数相加并打印结果
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片并打印其中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}
}