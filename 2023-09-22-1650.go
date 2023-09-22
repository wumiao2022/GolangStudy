package main

import (
	"fmt"
)

func main() {
	// 练习1: 定义一个结构体Person，包含姓名、年龄和性别三个字段，分别为字符串、整型和布尔型
	type Person struct {
		Name   string
		Age    int
		Gender bool
	}

	// 练习2: 创建一个名为"John"、年龄为28、性别为男的Person实例，并将其打印出来
	person1 := Person{
		Name:   "John",
		Age:    28,
		Gender: true,
	}

	fmt.Println(person1)

	// 练习3: 定义一个函数，接收一个Person类型的参数，并将其姓名和年龄打印出来
	printPerson := func(p Person) {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
	}

	printPerson(person1)
}