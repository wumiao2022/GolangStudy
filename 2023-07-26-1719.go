package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：变量使用
	var num1 int = 10
	var num2 int = 5
	var sum int = num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：循环
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习4：条件判断
	var score int = 80
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 练习5：函数调用
	result := add(3, 4)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}
