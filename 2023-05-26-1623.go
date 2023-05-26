package main

import "fmt"

func main() {
	// 练习题 1：打印 ASCII 码表
	for i := 0; i <= 127; i++ {
		fmt.Printf("%d: %c\n", i, i)
	}

	// 练习题 2：计算斐波那契数列的第 n 项
	n := 10
	fmt.Printf("斐波那契数列的第 %d 项是 %d\n", n, fibonacci(n))

	// 练习题 3：冒泡排序
	nums := []int{9, 6, 2, 8, 3, 1, 5, 4, 7}
	bubbleSort(nums)
	fmt.Println("排序后的结果为：", nums)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func bubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}