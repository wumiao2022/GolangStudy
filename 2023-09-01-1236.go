package main

import (
	"fmt"
)

func main() {
	// 练习一：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习二：计算两个数字的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习三：判断一个数是否为偶数
	num := 21
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习四：求一个数组的平均值
	numbers := []int{5, 10, 15, 20, 25}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	avg := float64(sum) / float64(len(numbers))
	fmt.Println("Average:", avg)
}