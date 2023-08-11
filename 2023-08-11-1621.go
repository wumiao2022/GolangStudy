package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印输出当前的时间
	fmt.Println(time.Now())

	// 2. 判断两个数是否相等
	num1 := 10
	num2 := 20
	fmt.Println(num1 == num2)

	// 3. 定义一个切片，包含整数1到10
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	// 4. 字符串连接
	str1 := "Hello"
	str2 := "Golang"
	fmt.Println(str1 + " " + str2)

	// 5. 判断字符串是否为空
	str := ""
	fmt.Println(len(str) == 0)

	// 6. 使用for循环输出1到10之间的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 7. 使用switch语句判断数字的奇偶性
	num := 5
	switch num % 2 {
	case 0:
		fmt.Println("偶数")
	default:
		fmt.Println("奇数")
	}

	// 8. 使用递归函数计算阶乘
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}