package main

import (
	"fmt"
)

func main() {
	// 练习1-输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2-交换变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println("a=", a, "b=", b)

	// 练习3-寻找两数之和等于目标值的下标
	nums := []int{2, 7, 11, 15}
	target := 9
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				fmt.Println(result)
				break
			}
		}
	}

	// 练习4-冒泡排序
	arr := []int{5, 2, 4, 6, 1, 3}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// 练习5-计算阶乘
	var fact uint64 = 1
	for i := 1; i <= 10; i++ {
		fact *= uint64(i)
	}
	fmt.Println("10的阶乘等于", fact)
}