package main

import "fmt"

func main() {
	// 实例1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 实例2：整数相加
	var num1, num2 int = 10, 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3：字符串拼接
	str1 := "Hello"
	str2 := "Golang"
	str := str1 + ", " + str2 + "!"
	fmt.Println(str)

	// 实例4：整数比较
	num3, num4 := 30, 40
	if num3 < num4 {
		fmt.Println("num3 is smaller than num4")
	} else {
		fmt.Println("num4 is smaller than num3")
	}

	// 实例5：数组遍历
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, "is", numbers[i])
	}

	// 实例6：函数调用
	result := add(5, 7)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}