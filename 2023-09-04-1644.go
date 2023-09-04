package main

import (
	"fmt"
)

func main() {
	// 1. 编写一个函数，接收两个整数参数并返回它们的和
	func sum(a, b int) int {
		return a + b
	}

	// 2. 编写一个函数，接收一个字符串参数，并返回该字符串的长度
	func stringLength(str string) int {
		return len(str)
	}

	// 3. 编写一个函数，接收一个整数参数，并判断该数字是否为偶数，是则返回true，否则返回false
	func isEven(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	}

	// 4. 编写一个函数，接收一个字符串切片参数，并返回该切片的元素个数
	func sliceLength(s []string) int {
		return len(s)
	}

	// 5. 编写一个函数，接收一个整数切片参数，并返回切片中所有元素的和
	func sumSlice(s []int) int {
		sum := 0
		for _, num := range s {
			sum += num
		}
		return sum
	}

	// 6. 编写一个函数，接收一个整数数组参数，并返回数组中的最大值和最小值
	func findMinMax(arr []int) (int, int) {
		min := arr[0]
		max := arr[0]
		for _, num := range arr {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		return max, min
	}

	// 7. 编写一个函数，接收一个整数参数n，并返回一个n行n列的乘法表字符串（每个元素之间使用制表符分隔）
	func multiplicationTable(n int) string {
		table := ""
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				table += fmt.Sprintf("%d\t", i*j)
			}
			table += "\n"
		}
		return table
	}
}