package main

import (
	"fmt"
)

func main() {
	// 练习1：将一个整数切片中的所有元素加1
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		numbers[i] += 1
	}
	fmt.Println(numbers)

	// 练习2：取一个字符串的前三个字符
	str := "hello"
	fmt.Println(str[:3])

	// 练习3：使用冒泡排序将一个整数切片进行升序排序
	numbers2 := []int{5, 2, 9, 1, 8}
	for i := 0; i < len(numbers2)-1; i++ {
		for j := 0; j < len(numbers2)-i-1; j++ {
			if numbers2[j] > numbers2[j+1] {
				numbers2[j], numbers2[j+1] = numbers2[j+1], numbers2[j]
			}
		}
	}
	fmt.Println(numbers2)
}