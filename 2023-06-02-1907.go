package main

import (
	"fmt"
)

func main() {
	// 1. 反转字符串
	str := "hello world"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 2. 查找最大值和最小值
	numbers := []int{3, 5, 1, 8, 9, 2}
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("Max: ", max)
	fmt.Println("Min: ", min)

	// 3. 计算斐波那契数列
	fib := []int{0, 1}
	for i := 0; i < 8; i++ {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}
	fmt.Println(fib)
}