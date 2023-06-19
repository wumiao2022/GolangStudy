package main

import (
	"fmt"
)

func main() {
	// 实现一个函数，用于将一个整数数组中的元素翻转
	intArr := []int{1, 2, 3, 4, 5}
	reverse(intArr)
	fmt.Println(intArr) // [5 4 3 2 1]

	// 实现一个递归函数，用于求解斐波那契数列的第 n 项
	n := 10
	fmt.Println(fibonacci(n)) // 55

	// 实现一个函数，用于对一个字符串数组进行排序
	strArr := []string{"banana", "apple", "orange"}
	sortStringArr(strArr)
	fmt.Println(strArr) // [apple banana orange]
}

// 翻转整数数组
func reverse(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
	}
}

// 递归实现斐波那契数列的第 n 项
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 对字符串数组进行排序
func sortStringArr(arr []string) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}