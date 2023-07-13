package main

import "fmt"

func main() {
	// 练习1: 定义一个函数，输入两个int类型的参数，返回它们的和
	func sum(a, b int) int {
		return a + b
	}

	fmt.Println(sum(5, 2))

	// 练习2: 定义一个函数，输入一个字符串，输出该字符串的长度
	func strLength(str string) int {
		return len(str)
	}

	fmt.Println(strLength("Hello World"))

	// 练习3: 定义一个函数，输入一个int类型的切片，输出该切片中的最大值
	func maxNumber(arr []int) int {
		max := arr[0]

		for _, num := range arr {
			if num > max {
				max = num
			}
		}

		return max
	}

	nums := []int{10, 5, 8, 3, 15}
	fmt.Println(maxNumber(nums))
}