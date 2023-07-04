package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2：使用循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：定义切片，并使用切片的append()函数追加元素
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	// 练习4：定义一个结构体，并初始化结构体对象的字段
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "Tom",
		Age:  25,
	}
	fmt.Println(person)
}

注意：以上代码仅为示例实例，具体实现方式可能根据实际需求有所不同。