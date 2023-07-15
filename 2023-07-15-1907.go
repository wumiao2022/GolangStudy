package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量声明与赋值
	var name string
	name = "John"
	age := 25
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// 练习3：控制流程 - if语句
	num := 10
	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// 练习4：循环 - for语句
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 10:", sum)

	// 练习5：数组和切片
	numbers := [5]int{1, 2, 3, 4, 5}
	slice := numbers[1:4]
	fmt.Println("Array:", numbers)
	fmt.Println("Slice:", slice)

	// 练习6：函数
	result := add(5, 3)
	fmt.Println("Result of add function:", result)
}

func add(a, b int) int {
	return a + b
}