package main

import "fmt"

func main() {
	// 示例1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 示例2：变量赋值与打印
	var name string
	name = "Alice"
	age := 25
	fmt.Println("My name is", name, "and I'm", age, "years old.")

	// 示例3：条件语句
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5")
	} else {
		fmt.Println("Number is less than or equal to 5")
	}

	// 示例4：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// 示例5：函数定义与调用
	result := calculateSum(10, 20)
	fmt.Println("The sum is:", result)
}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}