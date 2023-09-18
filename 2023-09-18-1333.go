package main

import "fmt"

func main() {
	// 实现一个函数，将一个数组中的元素按照逆序排列，并返回逆序后的数组
	reverseArray := func(arr []int) []int {
		n := len(arr)
		reversed := make([]int, n)
		for i := 0; i < n; i++ {
			reversed[i] = arr[n-i-1]
		}
		return reversed
	}

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseArray(arr)) // 输出 [5 4 3 2 1]
}
