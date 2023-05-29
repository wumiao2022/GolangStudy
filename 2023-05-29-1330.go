package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")
	
	// 2. 定义变量并赋初值，输出变量的值
	s := "Golang"
	fmt.Println(s)
	
	// 3. 判断数值的奇偶性
	num := 5
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	
	// 4. 循环输出数字1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 5. 定义一个slice并循环输出
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
	
	// 6. 使用函数计算两个数的和
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}