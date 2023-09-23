package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整型变量x，赋值为10，并打印出来
	x := 10
	fmt.Println(x)

	// 练习3：定义一个字符串变量name，赋值为"John"，并打印出来
	name := "John"
	fmt.Println(name)

	// 练习4：定义一个布尔型变量isTrue，并赋值为true，打印出来
	isTrue := true
	fmt.Println(isTrue)

	// 练习5：定义一个切片变量numbers，包含1、2、3三个元素，并打印出来
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// 练习6：定义一个函数add，接收两个整型参数，并返回它们的和
	result := add(5, 3)
	fmt.Println(result)

	// 练习7：定义一个结构体类型Person，包含姓名和年龄两个字段，并打印出一个Person对象的信息
	person := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
}