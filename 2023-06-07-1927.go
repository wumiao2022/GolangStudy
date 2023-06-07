package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello World!"
	fmt.Println("Hello World!")

	// 2. 输出 1-10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 创建一个切片，并依次输出其中的元素值
	s := []int{1, 2, 3, 4, 5}
	for _, value := range s {
		fmt.Println(value)
	}

	// 4. 定义一个函数，接收两个整数参数并返回他们的和
	add := func(a, b int) int {
		return a + b
	}
	result := add(1, 2)
	fmt.Println(result)

	// 5. 创建一个结构体，并实例化，设置其中字段的值，然后输出结构体的值
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Tom", Age: 18}
	fmt.Println(p)
}