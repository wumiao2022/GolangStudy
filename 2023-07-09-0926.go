package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 变量声明和赋值
	var x int
	x = 10
	fmt.Println("Value of x:", x)

	// 3. 算术运算
	a := 5
	b := 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 4. 条件语句
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	// 5. 循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 6. 数组
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 7. 切片
	slice := numbers[1:4]
	fmt.Println(slice)

	// 8. 函数
	result := add(3, 5)
	fmt.Println("Result of addition:", result)
}

// 函数：求两个数的和
func add(a, b int) int {
	return a + b
}