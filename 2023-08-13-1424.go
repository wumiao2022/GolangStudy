package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义和输出变量
	var num1 int = 10
	fmt.Println(num1)

	// 练习3: 数组的使用
	var nums [5]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5
	fmt.Println(nums)

	// 练习4: 循环和条件判断
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Println(nums[i], "is even")
		} else {
			fmt.Println(nums[i], "is odd")
		}
	}

	// 练习5: 使用函数
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// 练习6: 匿名函数和闭包
	addFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println("5 + 6 =", addFunc(5, 6))
}

// 函数的实现
func add(a, b int) int {
	return a + b
}