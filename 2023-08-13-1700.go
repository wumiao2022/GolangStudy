package main

import (
	"fmt"
)

func main() {
	// 1. 声明一个整型变量x并赋值为10
	x := 10
	fmt.Println(x)

	// 2. 声明一个字符串变量name并赋值为"Go语言"
	name := "Go语言"
	fmt.Println(name)

	// 3. 声明一个浮点型变量pi并赋值为3.1415
	pi := 3.1415
	fmt.Println(pi)

	// 4. 声明一个布尔型变量isTrue并赋值为true
	isTrue := true
	fmt.Println(isTrue)

	// 5. 声明一个整型切片numbers，并初始化为[1, 2, 3, 4, 5]
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 6. 声明一个map类型的变量person，并初始化为{"name": "张三", "age": 20, "gender": "男"}
	person := map[string]interface{}{
		"name":   "张三",
		"age":    20,
		"gender": "男",
	}
	fmt.Println(person)
}