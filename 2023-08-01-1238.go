package main

import (
	"fmt"
)

func main() {
	// 示例1：输出Hello, World!
	fmt.Println("Hello, World!")
	
	// 示例2：定义变量并打印
	var age int = 20
	fmt.Println("My age is", age)
	
	// 示例3：条件语句
	if age >= 18 {
		fmt.Println("I am an adult")
	} else {
		fmt.Println("I am a minor")
	}
	
	// 示例4：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("Current index:", i)
	}
	
	// 示例5：切片操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Length of nums:", len(nums))
	fmt.Println("First number in nums:", nums[0])
	
	// 示例6：函数
	fmt.Println("Sum of 3 and 4 is:", sum(3, 4))
}

func sum(a, b int) int {
	return a + b
}
