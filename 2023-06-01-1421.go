package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello world
	fmt.Println("Hello world")

	// 2. 定义变量并打印
	var name string = "Alice"
	fmt.Printf("My name is %s\n", name)

	// 3. 定义一个数组并打印
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 4. 定义一个map并打印
	m := map[string]string{"name": "Bob", "age": "20"}
	fmt.Println(m)

	// 5. 简单的if语句
	a, b := 1, 2
	if a > b {
		fmt.Println("a is bigger")
	} else {
		fmt.Println("b is bigger")
	}

	// 6. for循环打印1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 7. switch语句
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is not Monday or Tuesday")
	}

	// 8. 函数的定义和调用
	fmt.Println(add(3, 4))
}

func add(x, y int) int {
	return x + y
}