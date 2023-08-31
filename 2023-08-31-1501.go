package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和类型
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习3：条件语句
	num := 10
	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// 练习4：循环语句
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：函数
	fmt.Println(add(3, 4))

	// 练习6：切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers[2:4])

	// 练习7：结构体
	type Person struct {
		Name string
		Age  int
	}

	p := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(p.Name)

	// 练习8：指针
	x := 5
	updateValue(&x)
	fmt.Println(x)
}

func add(a, b int) int {
	return a + b
}

func updateValue(num *int) {
	*num = 10
}
