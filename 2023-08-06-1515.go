package main

import "fmt"

func main() {
	// 练习1：将两个整数相加并输出结果
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println(sum)

	// 练习2：定义一个字符串变量并输出
	str := "Hello, Golang!"
	fmt.Println(str)

	// 练习3：定义一个整数变量并判断是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 练习4：使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片并遍历输出
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}