package main

import (
	"fmt"
)

func main() {
	// 实例1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 实例2：求和函数
	sum := calculateSum(10, 20)
	fmt.Println("Sum:", sum)

	// 实例3：判断一个数是否为偶数
	num := 15
	if isEven(num) {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 实例4：取一个切片的前N个元素
	slice := []int{1, 2, 3, 4, 5}
	n := 3
	result := getFirstN(slice, n)
	fmt.Println("First", n, "elements:", result)
}

func calculateSum(a, b int) int {
	return a + b
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func getFirstN(slice []int, n int) []int {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}
