package main

import "fmt"

func main() {
	// 1. 打印Hello World
	fmt.Println("Hello World")

	// 2. 计算1+1并打印结果
	fmt.Println(1 + 1)

	// 3. 声明一个变量x，并将其赋值为10，打印变量的值
	x := 10
	fmt.Println(x)

	// 4. 声明一个字符串变量name，并将其赋值为你的名字，打印变量的值
	name := "John"
	fmt.Println(name)

	// 5. 声明一个整型数组a，包含1到5的数字，打印数组的长度和第3个元素的值
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a))
	fmt.Println(a[2])

	// 6. 使用for循环输出0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 7. 创建一个名为Person的结构体，包含name和age字段
	type Person struct {
		name string
		age  int
	}

	// 8. 创建一个Person类型的变量p，并将其name和age字段分别赋值为"Tom"和25，打印p的值
	p := Person{
		name: "Tom",
		age:  25,
	}
	fmt.Println(p)
}