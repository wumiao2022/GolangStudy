package main

import (
	"fmt"
)

// 练习1：交换两个变量的值
func swapVariables() {
	a := 10
	b := 20

	// 使用临时变量进行值交换
	temp := a
	a = b
	b = temp

	fmt.Println("交换后的值：")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

// 练习2：计算整数数组的和
func calculateSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

// 练习3：统计字符串中各个字符的频次
func countFrequency(str string) map[rune]int {
	frequencyMap := make(map[rune]int)
	for _, char := range str {
		frequencyMap[char]++
	}
	return frequencyMap
}

func main() {
	// 练习1：交换两个变量的值
	swapVariables()

	// 练习2：计算整数数组的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := calculateSum(numbers)
	fmt.Println("数组的和：", sum)

	// 练习3：统计字符串中各个字符的频次
	str := "Hello, World!"
	frequencyMap := countFrequency(str)
	fmt.Println("字符频次统计：")
	for char, frequency := range frequencyMap {
		fmt.Printf("%c: %d\n", char, frequency)
	}
}