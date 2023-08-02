package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量声明和赋值
	var num1 int = 10
	var num2 int = 20
	var sum int = num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：循环输出1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：切片(slice)操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("First Element:", nums[0])       // 获取切片第一个元素
	fmt.Println("Slice Length:", len(nums))       // 获取切片长度
	fmt.Println("Sliced Elements:", nums[1:3])   // 切片操作，获取第2到第3个元素（不包括第3个元素）

	// 练习5：函数定义和调用
	result := add(5, 3)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}