package main

import (
	"fmt"
)

func main() {
	// 练习1: 声明一个整型变量x，并赋予初始值为10
	var x int = 10
	fmt.Println(x)

	// 练习2: 声明一个浮点型变量y，并赋予初始值为3.14
	var y float64 = 3.14
	fmt.Println(y)

	// 练习3: 声明一个字符串变量z，并赋予初始值为"Hello, Golang!"
	var z string = "Hello, Golang!"
	fmt.Println(z)

	// 练习4: 使用短变量声明方式，声明一个整型变量a，并赋予初始值为5
	a := 5
	fmt.Println(a)

	// 练习5: 声明一个空切片b，长度为0
	var b []int
	fmt.Println(b)

	// 练习6: 声明一个包含5个元素的整型数组c，并赋予初始值为1, 2, 3, 4, 5
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	// 练习7: 使用循环遍历数组c，并将每个元素打印出来
	for _, num := range c {
		fmt.Println(num)
	}

	// 练习8: 声明一个函数add，接受两个整型参数a和b，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 练习9: 调用add函数，并将参数为2和3，将结果打印出来
	result := add(2, 3)
	fmt.Println(result)

	// 练习10: 声明一个结构体person，包含姓名和年龄字段，并赋予初始值
	type person struct {
		name string
		age  int
	}

	p := person{
		name: "Alice",
		age:  28,
	}

	fmt.Println(p)
}