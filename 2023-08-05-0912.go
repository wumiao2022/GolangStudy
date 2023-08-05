package main

import "fmt"

func main() {
	// 练习1: 输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数并输出结果
	number := 13
	if number%2 == 0 {
		fmt.Println("This number is even.")
	} else {
		fmt.Println("This number is odd.")
	}

	// 练习4: 循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 使用switch语句根据成绩输出对应的等级
	score := 80
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B")
	case score >= 70:
		fmt.Println("Grade C")
	case score >= 60:
		fmt.Println("Grade D")
	default:
		fmt.Println("Grade F")
	}
}