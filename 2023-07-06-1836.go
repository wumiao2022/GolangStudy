package main

import "fmt"

func main() {
	// 练习1：打印字符串
	fmt.Println("Hello, world!")

	// 练习2：整数相加
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断奇偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：循环打印数字
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 练习5：切片操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original nums:", nums)
	fmt.Println("First element:", nums[0])
	fmt.Println("Last element:", nums[len(nums)-1])
	fmt.Println("Sliced nums:", nums[1:3])
	fmt.Println("Appended nums:", append(nums, 6))
}