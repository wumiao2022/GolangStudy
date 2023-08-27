package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个整数的和
	num1 := 7
	num2 := 9
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为奇数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习5：使用数组存储五个整数，并打印出来
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}

	// 练习6：使用切片存储五个整数，并打印出来
	numbersSlice := []int{6, 7, 8, 9, 10}
	for _, number := range numbersSlice {
		fmt.Println(number)
	}

	// 练习7：使用map存储三个学生的成绩，并打印出来
	scores := map[string]int{
		"Alice":  80,
		"Bob":    90,
		"Cynthia": 85,
	}
	for name, score := range scores {
		fmt.Println(name, ":", score)
	}

	// 练习8：使用函数计算两个数的乘积
	product := multiply(4, 5)
	fmt.Println("Product:", product)
}

func multiply(num1, num2 int) int {
	return num1 * num2
}