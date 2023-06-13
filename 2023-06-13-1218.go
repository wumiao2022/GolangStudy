package main

import (
	"fmt"
)

func main() {
	// 1. 声明变量和常量
	var a int = 10
	const b = 5

	// 2. 条件语句（if else）
	if a > b {
		fmt.Println("a 大于 b")
	} else {
		fmt.Println("a 小于等于 b")
	}

	// 3. 循环语句（for）
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 4. 数组
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	// 5. 切片
	slice := make([]int, 5)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Println(slice)

	// 6. 映射
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// 7. 函数
	fmt.Println(add(1, 2))

	// 8. 结构体
	p := person{
		name: "张三",
		age:  30,
	}
	fmt.Println(p)
}

func add(a, b int) int {
	return a + b
}

type person struct {
	name string
	age  int
}