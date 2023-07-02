package main

import "fmt"

func main() {
	// 实例1：Hello World
	fmt.Println("Hello, World!")

	// 实例2：变量和常量
	var name string = "GoLang"
	const age int = 10

	// 实例3：循环和条件语句
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 实例4：函数
	result := sum(5, 3)
	fmt.Println("Sum:", result)

	// 实例5：结构体和方法
	p := Person{name: "Alice", age: 25}
	p.Introduce()
}

func sum(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}

func (p Person) Introduce() {
	fmt.Printf("My name is %s and my age is %d\n", p.name, p.age)
}
