package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数
	num := 30
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 循环打印数字1到100
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// 练习5: 定义一个结构体，并输出其字段值
	type person struct {
		name string
		age  int
	}
	p := person{name: "Alice", age: 25}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}