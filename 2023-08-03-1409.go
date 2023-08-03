package main

import "fmt"

func main() {
	// 练习1: 输出Hello World
	fmt.Println("Hello World")

	// 练习2: 变量与常量
	var num1 int = 10
	num2 := 5
	const pi float64 = 3.14159

	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("pi * num1 =", pi*num1)

	// 练习3: 条件语句
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	// 练习4: 循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("Current number:", i)
	}

	// 练习5: 切片与循环遍历
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习6: 函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 练习7: 结构体与方法
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("Area of rectangle:", rectangle.getArea())
}

// 函数示例
func add(a, b int) int {
	return a + b
}

// 结构体示例
type Rectangle struct {
	width  int
	height int
}

// 方法示例
func (r Rectangle) getArea() int {
	return r.width * r.height
}