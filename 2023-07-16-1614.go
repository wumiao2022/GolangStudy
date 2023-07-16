package main

import (
	"fmt"
)

func main() {
	// 1. 使用for循环打印出0-9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2. 打印出1-20之间的偶数
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 3. 判断一个数字是否是素数（只能被1和自身整除）
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("是素数")
	} else {
		fmt.Println("不是素数")
	}

	// 4. 打印出斐波那契数列的前n项（斐波那契数列的每一项都是前两项之和）
	n := 10
	prev := 0
	curr := 1
	for i := 0; i < n; i++ {
		fmt.Println(curr)
		temp := curr
		curr = prev + curr
		prev = temp
	}
}

// 5. 实现一个函数，接收一个字符串作为参数，返回该字符串的反转结果
func reverse(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}