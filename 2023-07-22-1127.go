package main

import (
	"fmt"
)

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 变量定义和使用
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I'm %d years old.\n", name, age)

	// 3. 数组和遍历
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 4. 循环判断
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even.\n", i)
		} else {
			fmt.Printf("%d is odd.\n", i)
		}
	}

	// 5. 函数定义和调用
	result := sum(10, 20)
	fmt.Println("Sum:", result)
}

func sum(a, b int) int {
	return a + b
}