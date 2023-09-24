package main

import "fmt"

func main() {
	// 1. 输出Hello, world!
	fmt.Println("Hello, world!")

	// 2. 定义一个整型变量x并赋值为10，打印出x的值
	x := 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name并赋值为"John"，打印出name的值
	name := "John"
	fmt.Println(name)

	// 4. 创建一个整型数组a，包含元素1, 2, 3, 4, 5，遍历数组并打印出每个元素的值
	a := [5]int{1, 2, 3, 4, 5}
	for _, num := range a {
		fmt.Println(num)
	}

	// 5. 创建一个map，包含键值对{name: "John", age: 30}，打印出map的内容
	person := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	fmt.Println(person)
}