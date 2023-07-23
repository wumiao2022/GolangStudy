package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用 for 循环输出 1~10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，将两个数相加并返回结果
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Result of addition:", add(5, 3))

	// 练习6：定义一个结构体，表示一个人的信息，并输出其中的字段值
	type person struct {
		name string
		age  int
	}
	p := person{name: "John", age: 25}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}