package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数之和，并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数，并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：使用for循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用switch语句判断一个人的年龄段，并打印结果
	age := 25
	switch {
	case age < 18:
		fmt.Println("Underage")
	case age >= 18 && age < 35:
		fmt.Println("Young adult")
	default:
		fmt.Println("Adult")
	}
}