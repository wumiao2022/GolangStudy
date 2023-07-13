package main

import "fmt"

func main() {
	// 例一：计算两个数的和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 例二：判断一个数是奇数还是偶数并输出结果
	num := 3
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 例三：判断一个年份是否是闰年并输出结果
	year := 2020
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}

	// 例四：使用循环打印出斐波那契数列的前20个数
	fmt.Println("Fibonacci series:")
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}

	// 例五：创建一个包含10个元素的整型数组，并遍历输出数组的元素
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array elements:")
	for _, num := range arr {
		fmt.Println(num)
	}
}