package main

import "fmt"

func main() {
	// 练习1：打印HelloWorld
	fmt.Println("Hello, World!")

	// 练习2：变量声明和使用
	var name string = "John"
	var age int = 25
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")

	// 练习3：数组和循环
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 练习4：切片和迭代
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	for index, value := range names {
		fmt.Println(index, value)
	}

	// 练习5：函数和参数
	result := add(3, 5)
	fmt.Println("The sum is", result)
}

func add(a, b int) int {
	return a + b
}