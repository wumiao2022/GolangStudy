package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和常量
	var x int = 10
	const pi float64 = 3.14
	
	// 练习3：条件语句
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// 练习4：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 练习5：数组和切片
	arr := [3]int{1, 2, 3}
	slice := []string{"apple", "banana", "orange"}
	
	// 练习6：函数
	result := add(2, 3)
	fmt.Println("Sum:", result)
}

func add(a, b int) int {
	return a + b
}