package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 定义一个整型变量，并将其赋值为10，然后打印输出
	var num int = 10
	fmt.Println(num)

	// 3. 定义一个字符串变量，并将其赋值为"Go语言练习"
	var str string = "Go语言练习"
	fmt.Println(str)

	// 4. 定义一个数组，包含元素1, 2, 3, 4, 5，然后打印输出
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 5. 定义一个切片，包含元素6, 7, 8，然后打印输出
	slice := []int{6, 7, 8}
	fmt.Println(slice)

	// 6. 定义一个map，键为字符串类型，值为整型，包含键值对"one": 1, "two": 2, "three": 3，然后打印输出
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	// 7. 定义一个结构体，包含字段name和age，并将其实例化，然后打印输出
	type Person struct {
		name string
		age  int
	}
	p := Person{name: "Tony", age: 30}
	fmt.Println(p)
}