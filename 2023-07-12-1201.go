package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 输出字符串 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 定义一个整型变量x并赋值为42，打印该变量的值
	x := 42
	fmt.Println(x)

	// 3. 定义一个字符串变量name并赋值为"John"，打印该变量的值
	name := "John"
	fmt.Println(name)

	// 4. 定义一个浮点型变量pi并赋值为3.14159，打印该变量的值
	pi := 3.14159
	fmt.Println(pi)

	// 5. 定义一个布尔型变量isTrue并赋值为true，打印该变量的值
	isTrue := true
	fmt.Println(isTrue)

	// 6. 使用for循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 7. 获取当前日期和时间，并打印出来
	currentTime := time.Now()
	fmt.Println(currentTime)
}