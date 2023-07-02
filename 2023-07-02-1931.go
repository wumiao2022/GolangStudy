package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和常量
	var name string = "John"
	age := 25
	const country string = "USA"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)

	// 练习3：数组和切片
	numbers := [5]int{1, 2, 3, 4, 5}
	slice := numbers[1:4]

	fmt.Println("Array:", numbers)
	fmt.Println("Slice:", slice)

	// 练习4：循环和条件语句
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 练习5：函数
	result := add(10, 20)
	fmt.Println("Result of add function:", result)

	// 练习6：指针
	num := 10
	ptr := &num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)

	// 练习7：结构体和方法
	rect := Rectangle{width: 10, height: 5}
	area := rect.calculateArea()
	fmt.Println("Area of rectangle:", area)
}

func add(a, b int) int {
	return a + b
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) calculateArea() int {
	return r.width * r.height
}