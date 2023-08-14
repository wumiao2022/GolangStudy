package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算并打印1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sum)

	// 练习4：定义一个结构体Person，并为其添加一个方法PrintName，可以打印出人的姓名
	type Person struct {
		Name string
	}
	
	p := Person{Name: "John Doe"}
	p.PrintName()
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}