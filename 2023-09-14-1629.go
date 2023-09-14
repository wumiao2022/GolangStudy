package main

import "fmt"

func main() {
	// 实例1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 实例2: 定义和使用变量
	var name string = "John"
	age := 30
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// 实例3: 数组和循环
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 实例4: 函数和多返回值
	result := sum(5, 10)
	fmt.Println("The sum is", result)

	// 实例5: 结构体和方法
	p := Person{name: "Alice", age: 25}
	p.SayHello()
}

func sum(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}