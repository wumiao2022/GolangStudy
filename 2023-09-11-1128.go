package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 声明一个整型变量，并打印出来
	var num int = 10
	fmt.Println(num)

	// 2. 声明一个字符串变量，并打印出来
	var str string = "Hello, Golang!"
	fmt.Println(str)

	// 3. 声明一个布尔变量，并打印出来
	var flag bool = true
	fmt.Println(flag)

	// 4. 使用for循环打印出1到10之间的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 获取当前的时间，并打印出来
	now := time.Now()
	fmt.Println(now)
}