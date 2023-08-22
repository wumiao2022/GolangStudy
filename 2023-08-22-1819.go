package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 练习4：定义一个结构体并创建实例
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Alice", 25}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}