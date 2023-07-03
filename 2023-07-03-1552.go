package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习3：判断一个数是否为偶数
	number := 13
	if number%2 == 0 {
		fmt.Println(number, "is an even number.")
	} else {
		fmt.Println(number, "is an odd number.")
	}

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用switch语句判断一个人的年龄段
	age := 25
	switch {
	case age < 18:
		fmt.Println("Underage")
	case age >= 18 && age < 30:
		fmt.Println("Young adult")
	default:
		fmt.Println("Adult")
	}
}