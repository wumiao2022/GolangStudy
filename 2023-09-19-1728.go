package main

import "fmt"

func main() {
	// 示例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 示例2：计算两个数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例3：遍历切片并打印元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 示例4：使用if-else语句判断数值大小并打印结果
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is smaller than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}
}