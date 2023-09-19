package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前日期
	fmt.Println(time.Now().Format("2006-01-02"))

	// 定义一个整数切片，对切片进行排序，然后打印排序后的结果
	numbers := []int{9, 3, 6, 1, 8, 2, 5, 7, 4}
	sort(numbers)
	fmt.Println(numbers)

	// 编写一个函数，接受一个字符串作为参数，并返回将字符串中的每个字符翻倍的结果
	doubled := doubleChars("hello")
	fmt.Println(doubled)
}

// 冒泡排序算法
func sort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

// 将字符串中的每个字符翻倍
func doubleChars(str string) string {
	result := ""
	for _, c := range str {
		result += string(c) + string(c)
	}
	return result
}