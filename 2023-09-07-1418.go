package main

import (
	"fmt"
)

func main() {
	// 实例1：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 实例2：判断一个数字是否为偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("这个数字是偶数")
	} else {
		fmt.Println("这个数字是奇数")
	}

	// 实例3：计算斐波那契数列的前n个数字
	n := 10
	first := 0
	second := 1
	fmt.Println(first)
	fmt.Println(second)
	for i := 3; i <= n; i++ {
		next := first + second
		fmt.Println(next)
		first = second
		second = next
	}
}