package main

import "fmt"

func main() {
	// 练习1：声明一个整数数组，包含5个元素，并赋予任意值，然后打印出数组中的所有元素

	numbers := [5]int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习2：声明一个字符串变量，赋值为 "Hello, World!"，然后打印出该字符串的长度和第一个字母

	str := "Hello, World!"
	fmt.Println("Length of string:", len(str))
	fmt.Println("First letter:", string(str[0]))

	// 练习3：声明一个函数，接收两个整数参数，并返回它们的和作为结果

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Result of addition:", add(10, 20))

	// 练习4：声明一个结构体类型，包含姓名和年龄两个字段，然后创建一个结构体实例并打印出其字段值

	type Person struct {
		Name string
		Age  int
	}

	p := Person{"John Doe", 30}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}