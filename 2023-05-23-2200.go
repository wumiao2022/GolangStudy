package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello world!"
	fmt.Println("Hello world!")

	// 2. 声明一个名为 name 的字符串变量，赋值为 "Golang"
	name := "Golang"

	// 3. 打印出字符串变量 name 的值
	fmt.Println(name)

	// 4. 声明一个名为 num 的整数变量，赋值为 42
	num := 42

	// 5. 打印出整数变量 num 的值
	fmt.Println(num)

	// 6. 声明一个名为 isRaining 的布尔变量，赋值为 true
	isRaining := true

	// 7. 打印出布尔变量 isRaining 的值
	fmt.Println(isRaining)

	// 8. 声明一个名为 nums 的整数切片，赋值为 1、2、3
	nums := []int{1, 2, 3}

	// 9. 打印出整数切片 nums 的值
	fmt.Println(nums)

	// 10. 使用 for 循环打印出整数切片 nums 中的每一个值
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 11. 使用 range 关键字打印出整数切片 nums 中的每一个值
	for _, num := range nums {
		fmt.Println(num)
	}
}