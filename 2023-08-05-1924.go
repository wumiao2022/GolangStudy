package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并输出
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：比较两个数的大小并输出较大的数
	if num1 > num2 {
		fmt.Println("较大的数是:", num1)
	} else {
		fmt.Println("较大的数是:", num2)
	}

	// 练习4：计算一个数的平方并输出
	num3 := 7
	square := num3 * num3
	fmt.Println("平方:", square)

	// 练习5：判断一个数是否为偶数并输出结果
	num4 := 12
	if num4%2 == 0 {
		fmt.Println(num4, "是偶数")
	} else {
		fmt.Println(num4, "不是偶数")
	}
}