package main

import (
	"fmt"
)

func main() {
	// 1. 变量和常量
	var a int = 10
	b := 20
	const c string = "hello world"

	// 2. 数组和切片
	arr := [3]int{1, 2, 3}
	slice := []string{"apple", "banana", "orange"}

	// 3. 控制结构
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range slice {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	// 4. 函数
	sum := add(1, 2)
	fmt.Println(sum)

	// 5. 接口和类型断言
	var varInterface interface{} = "hello"
	fmt.Println(varInterface.(string))

	// 6. 并发
	channel := make(chan int)
	go printNum(channel)
	fmt.Println(<-channel)
}

func add(a, b int) int {
	return a + b
}

func printNum(channel chan int) {
	channel <- 10
}