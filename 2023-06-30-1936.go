package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和并输出
	var num1, num2 int = 10, 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并输出结果
	num := 33
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组存储5个整数，并输出它们的和
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Sum of array:", sum)
}
