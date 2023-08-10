package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量赋值与打印
	name := "Alice"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习3：条件语句
	num := 10
	if num > 5 {
		fmt.Println("The number is greater than 5.")
	} else {
		fmt.Println("The number is not greater than 5.")
	}

	// 练习4：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 练习5：切片操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[:3]) // 输出 [1 2 3]

	// 练习6：函数定义与调用
	result := add(3, 4)
	fmt.Println(result) // 输出 7
}

func add(a, b int) int {
	return a + b
}