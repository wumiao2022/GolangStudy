package main

import (
	"fmt"
)

func main() {
	// 1. 变量声明和赋值
	var a int = 10
	b := 20
	fmt.Println(a, b)

	// 2. 数组和切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(arr, slice)

	// 3. 循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 4. 函数
	sum := add(1, 2)
	fmt.Println(sum)
}

func add(x, y int) int {
	return x + y
}