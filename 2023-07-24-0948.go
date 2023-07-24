package main

import "fmt"

func main() {
	// 练习题1：编写一个函数，实现两个整数相加并返回结果
	func sum(a int, b int) int {
		return a + b
	}

	// 练习题2：编写一个函数，实现判断一个数是否为偶数，并返回布尔值
	func isEven(num int) bool {
		if num%2 == 0 {
			return true
		} else {
			return false
		}
	}

	// 练习题3：编写一个函数，实现计算一个数组中所有元素的和，并返回结果
	func sumArray(arr []int) int {
		sum := 0
		for _, num := range arr {
			sum += num
		}
		return sum
	}

	// 练习题4：编写一个函数，实现找到一个数组中最大的元素，并返回结果
	func findMax(arr []int) int {
		max := arr[0]
		for _, num := range arr {
			if num > max {
				max = num
			}
		}
		return max
	}

	fmt.Println(sum(2, 3))
	fmt.Println(isEven(4))
	fmt.Println(sumArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(findMax([]int{10, 5, 8, 3, 15}))
}