package main

import (
	"fmt"
)

func main() {
	// 实现字符串反转函数
	fmt.Println(reverseString("hello world"))

	// 实现斐波那契数列函数
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	// 实现冒泡排序算法
	arr := []int{5, 2, 4, 6, 1, 3}
	bubbleSort(arr)
	fmt.Println(arr)

	// 实现快速排序算法
	arr2 := []int{5, 2, 4, 6, 1, 3}
	quickSort(arr2)
	fmt.Println(arr2)
}

// 字符串反转函数
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 斐波那契数列函数
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 冒泡排序算法
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 快速排序算法
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[0]
	left, right := 1, len(arr)-1
	for left <= right {
		if arr[left] <= pivot {
			left++
			continue
		}
		if arr[right] > pivot {
			right--
			continue
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[0], arr[right] = arr[right], arr[0]
	quickSort(arr[:right])
	quickSort(arr[right+1:])
}