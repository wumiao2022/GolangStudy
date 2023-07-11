package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：变量声明和赋值
	var name string = "John"
	var age int = 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 练习3：循环打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：使用条件语句判断数字是否为偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习5：数组和切片操作
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[2:4]
	slice[0] = 7
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// 练习6：函数的定义和调用
	result := add(3, 4)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}