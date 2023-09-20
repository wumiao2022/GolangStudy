package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum of", num1, "and", num2, "is", sum)

	// 练习3：输出1-10之间的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：使用循环计算1-10之间的整数的阶乘，并打印出结果
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
		fmt.Println("Factorial of", i, "is", factorial)
	}

	// 练习5：定义一个结构体Person，包含name和age两个字段，并创建一个Person对象并打印出来
	type Person struct {
		name string
		age  int
	}

	person := Person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(person)
}