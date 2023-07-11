package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数相加的结果，并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数，并输出判断结果
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习4：使用for循环输出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组存储5个整数，并求平均值
	numbers := [5]int{1, 2, 3, 4, 5}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	avg := float64(sum) / float64(len(numbers))
	fmt.Println("Average:", avg)
}