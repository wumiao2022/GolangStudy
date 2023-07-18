package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出当前时间
	fmt.Println(time.Now())

	// 练习2：定义一个整型变量，将其初始化为10，然后打印出变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：定义一个整型数组，将其中的元素依次赋值为1，2，3，4，5，然后打印出数组的内容
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习4：定义一个map，其中key为string类型，value为int类型，将其中的元素依
	// 次赋值为"one": 1, "two": 2, "three": 3, 然后打印出map的内容
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	// 练习5：定义一个函数，接收一个整型参数n，返回n的平方值
	fmt.Println(square(5))

	// 练习6：定义一个结构体类型Person，包含姓名和年龄字段，
	// 然后创建一个Person实例，赋值并打印出该实例
	p := Person{
		Name: "Tom",
		Age:  25,
	}
	fmt.Println(p)
}

func square(n int) int {
	return n * n
}

type Person struct {
	Name string
	Age  int
}