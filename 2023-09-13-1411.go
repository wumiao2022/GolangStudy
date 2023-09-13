package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 2. 定义一个整数变量x，并打印其值
	var x int = 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name，并打印其值
	var name string = "Alice"
	fmt.Println(name)
	
	// 4. 计算两个整数的和，并打印结果
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println(sum)
	
	// 5. 创建一个切片，并打印其长度和容量
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	
	// 6. 编写一个函数，将两个整数作为参数，返回它们的乘积
	result := multiply(4, 6)
	fmt.Println(result)
	
	// 7. 使用循环打印出1到10之间的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func multiply(a, b int) int {
	return a * b
}