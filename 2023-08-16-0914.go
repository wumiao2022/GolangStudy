package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出Hello World!
	fmt.Println("Hello World!")

	// 练习2: 变量赋值及打印
	name := "Gopher"
	age := 35
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习3: 切片操作
	nums := []int{1, 2, 3, 4, 5}
	firstTwo := nums[:2]
	lastThree := nums[2:]
	fmt.Println("First two numbers:", firstTwo)
	fmt.Println("Last three numbers:", lastThree)

	// 练习4: 循环
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习5: 函数
	sum := add(2, 3)
	fmt.Println("Sum:", sum)

	// 练习6: 结构体
	person := Person{name: "Alice", age: 25}
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}