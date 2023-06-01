package main

import (
	"fmt"
)

func main() {
	// 1. 输出"Hello, Go!"。
	fmt.Println("Hello, Go!")

	// 2. 定义一个变量x，赋值为10，将其输出。
	x := 10
	fmt.Println(x)

	// 3. 使用for循环输出1到10的数字。
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 4. 定义一个切片a，包含元素1、2、3、4、5，将其输出。
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// 5. 定义一个空map，添加键值对"apple": 1，"banana": 2，将其输出。
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	fmt.Println(m)
}