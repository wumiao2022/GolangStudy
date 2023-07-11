package main

import (
	"fmt"
)

func main() {
	// 1. 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 2. 声明并初始化一个字符串变量，输出该字符串的长度
	str := "Hello, World!"
	fmt.Println(len(str))

	// 3. 声明一个整型数组，并遍历数组的每个元素并输出
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 4. 声明一个map，包含学生姓名和对应的年龄信息，输出map的所有键值对
	students := map[string]int{
		"Alice":  20,
		"Bob":    22,
		"Charlie": 18,
	}
	for name, age := range students {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// 5. 声明一个切片，包含5个整数，使用冒泡排序算法将切片元素从小到大排序，并输出排序后的结果
	nums := []int{9, 5, 7, 3, 1}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}