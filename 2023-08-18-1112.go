package main

import "fmt"

func main() {
	// 1. 声明一个名为count的整型变量，并赋值为10
	count := 10
	fmt.Println(count)

	// 2. 使用循环打印出1~10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 声明一个名为message的字符串变量，并赋值为"Hello, Golang!"
	message := "Hello, Golang!"
	fmt.Println(message)

	// 4. 声明一个名为numbers的整型切片，包含数字1到5
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 5. 声明一个名为person的map，包含键值对name: "John"和age: 30
	person := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	fmt.Println(person)
}