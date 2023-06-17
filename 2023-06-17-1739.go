package main

import "fmt"

func main() {
	// 练习1：输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义变量并输出
	var name string = "Gopher"
	fmt.Printf("My name is %s\n", name)

	// 练习3：定义函数并调用
	result := add(3, 5)
	fmt.Printf("The result is %d\n", result)

	// 练习4：定义结构体并使用
	p := person{name: "Alice", age: 25}
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)

	// 练习5：循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习6：条件语句
	score := 85
	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func add(x, y int) int {
	return x + y
}

type person struct {
	name string
	age  int
}