package main

import "fmt"

func main() {
	// 1. 打印出Hello, world!
	fmt.Println("Hello, world!")

	// 2. 声明一个字符串变量并打印出来
	message := "Hello, Golang!"
	fmt.Println(message)

	// 3. 声明一个整数变量并打印出来
	number := 100
	fmt.Println(number)

	// 4. 声明一个布尔变量并打印出来
	isTrue := true
	fmt.Println(isTrue)

	// 5. 声明一个浮点数变量并打印出来
	pi := 3.14159
	fmt.Println(pi)

	// 6. 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 7. 使用if条件判断语句判断一个整数是正数、负数还是零，并打印出结果
	number := -10
	if number > 0 {
		fmt.Println("Positive")
	} else if number < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// 8. 声明一个包含5个字符串元素的数组，并使用for循环将这些元素打印出来
	strings := [5]string{"apple", "banana", "orange", "pear", "grape"}
	for _, item := range strings {
		fmt.Println(item)
	}

	// 9. 使用函数实现两个整数相加并返回结果
	result := add(10, 20)
	fmt.Println(result)

	// 10. 使用结构体来表示一个人的信息，并将其打印出来
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}
	fmt.Printf("Name: %s, Age: %d, Country: %s\n", person.Name, person.Age, person.Country)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	Name    string
	Age     int
	Country string
}