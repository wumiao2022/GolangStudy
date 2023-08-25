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

	// 练习3：判断一个数是否为偶数
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 练习4：使用循环输出10以内的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	//  练习5：使用数组存储和输出5个整数
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}