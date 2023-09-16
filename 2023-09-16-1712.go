package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// 练习1: 遍历并打印出 numbers 切片中的所有元素
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习2: 使用 for 循环计算并打印出 numbers 切片中所有元素的和
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 练习3: 使用 for 循环找出 numbers 切片中的最大值，并打印出来
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max:", max)

	// 练习4: 使用 for 循环找出 numbers 切片中的最小值，并打印出来
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	fmt.Println("Min:", min)
}