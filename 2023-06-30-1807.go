package main

import "fmt"

func main() {
	// 1. 打印出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 定义一个函数，接收两个整数参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))

	// 3. 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 4. 使用if-else语句判断一个数是奇数还是偶数，并打印结果
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 5. 定义一个结构体类型Person，包含姓名和年龄字段，创建一个Person实例并打印出来
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
	    Name: "John",
	    Age:  25,
	}
	fmt.Println(person)
}