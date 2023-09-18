package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World!")

	// 练习2: 计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为奇数
	num := 17
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 定义并使用一个结构体
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
		Name: "John",
		Age:  25,
	}
	fmt.Println(person)

	// 练习6: 使用切片进行元素操作
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	fmt.Println(numbers)

	// 练习7: 使用map存储键值对
	studentGrades := make(map[string]int)
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	fmt.Println(studentGrades)

	// 练习8: 使用函数使用递归计算阶乘
	n := 5
	factorial := calculateFactorial(n)
	fmt.Println("Factorial of", n, "is", factorial)

	// 练习9: 使用接口实现多态
	var shape Shape
	shape = Circle{radius: 5}
	fmt.Println("Area of Circle:", shape.calculateArea())
	shape = Rectangle{width: 10, height: 8}
	fmt.Println("Area of Rectangle:", shape.calculateArea())
}

type Shape interface {
	calculateArea() float64
}

type Circle struct {
	radius float64
}

func (c Circle) calculateArea() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) calculateArea() float64 {
	return r.width * r.height
}

func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}