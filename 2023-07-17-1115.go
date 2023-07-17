package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出Hello World
	fmt.Println("Hello World!")

	// 练习2: 定义变量并输出
	var name string = "Alice"
	fmt.Println("My name is", name)

	// 练习3: 通过循环输出数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 使用条件语句判断年龄段
	age := 25
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// 练习5: 定义一个数组并遍历输出
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 练习6: 定义一个结构体并打印字段值
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Bob", 30}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}