package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：变量声明与赋值
	var num1 int = 10
	num2 := 20
	fmt.Println(num1 + num2)

	// 练习3：条件语句
	num := 5
	if num > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}

	// 练习4：循环语句
	count := 0
	for i := 1; i <= 10; i++ {
		count += i
	}
	fmt.Println(count)

	// 练习5：切片操作
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[2:4])

	// 练习6：函数定义与调用
	result := add(10, 20)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}
```