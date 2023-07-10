package main

import (
	"fmt"
)

func main() {
	// 实例1：计算两个整数的和并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例2：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 实例3：交换两个变量的值
	a := 6
	b := 3
	fmt.Println("Before swap: a =", a, "b =", b)
	a, b = b, a
	fmt.Println("After swap: a =", a, "b =", b)

	// 实例4：数组遍历并输出元素值之和
	numbers := []int{2, 4, 6, 8, 10}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of numbers:", sum)
}