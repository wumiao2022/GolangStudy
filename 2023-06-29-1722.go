package main

import (
	"fmt"
)

func main() {
	// 实现一个函数，接收一个整数切片，并返回切片中的最小值和最大值
	func findMinMax(numbers []int) (int, int) {
		min := numbers[0]
		max := numbers[0]

		for _, num := range numbers {
			if num < min {
				min = num
			}

			if num > max {
				max = num
			}
		}

		return min, max
	}

	numbers := []int{15, 3, 9, 7, 11, 5}
	min, max := findMinMax(numbers)
	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
}