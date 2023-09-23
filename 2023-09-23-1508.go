package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 声明并初始化一个整数变量，然后打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3: 声明一个字符串变量，并将其赋值为你的名字，然后打印该变量的值
	name := "John"
	fmt.Println(name)

	// 练习4: 声明并初始化一个布尔变量，并将其赋值为真，然后打印该变量的值
	isTrue := true
	fmt.Println(isTrue)

	// 练习5: 声明一个空的切片变量，然后向其添加三个整数元素，并打印该切片的长度
	var numbers []int
	numbers = append(numbers, 1, 2, 3)
	fmt.Println(len(numbers))

	// 练习6: 声明一个map变量，然后向其添加三个键值对，并打印该map的值
	person := make(map[string]string)
	person["name"] = "Alice"
	person["age"] = "30"
	person["city"] = "New York"
	fmt.Println(person)
}