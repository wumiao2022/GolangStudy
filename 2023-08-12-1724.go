package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前日期
	fmt.Println("当前日期:", time.Now().Format("2006-01-02"))

	// 声明一个切片
	numbers := []int{1, 2, 3, 4, 5}

	// 遍历切片并打印元素
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 声明一个map并初始化
	student := map[string]int{
		"Tom":   90,
		"Jerry": 80,
		"Amy":   95,
	}

	// 遍历map并打印键值对
	for name, score := range student {
		fmt.Printf("%s的分数是%d\n", name, score)
	}

	// 声明一个结构体
	type Person struct {
		Name string
		Age  int
	}

	// 初始化结构体
	person := Person{
		Name: "John",
		Age:  25,
	}

	// 打印结构体字段
	fmt.Printf("姓名：%s，年龄：%d\n", person.Name, person.Age)
}