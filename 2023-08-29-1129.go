package main

import "fmt"

func main() {
	// 实例1：打印Hello World
	fmt.Println("Hello World")

	// 实例2：变量赋值和打印
	name := "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	// 实例3：算术运算
	num1 := 10
	num2 := 5
	sum := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num1 / num2
	fmt.Printf("Sum: %d, Subtraction: %d, Multiplication: %d, Division: %d\n", sum, subtraction, multiplication, division)

	// 实例4：判断语句
	score := 80
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// 实例5：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 实例6：函数
	result := add(2, 3)
	fmt.Println(result)
}

func add(num1, num2 int) int {
	return num1 + num2
}
