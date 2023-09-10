package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")

	// 1. 输出"Hello, Go!"到控制台

	// 2. 声明一个整型变量x，并将其初始化为10

	// 3. 声明一个字符串变量name，将其初始化为"Go Programming"

	// 4. 使用循环打印出x的值，直到x的值大于20为止

	// 5. 使用判断语句判断name是否等于"Go Programming"，如果是则输出"Hello, Golang!"，否则输出"Hello, World!"

	// 6. 声明一个切片slice，包含整数1、2、3、4、5

	// 7. 使用for range循环遍历slice，并打印出每一个元素的值

	// 8. 声明一个map类型m，键为string类型，值为int类型

	// 9. 向map m中添加键值对，键为"apple"，值为10

	// 10. 使用fmt.Printf打印出map m的键值对

	// 11. 声明一个函数add，接收两个整型参数a和b，返回它们的和

	// 12. 调用函数add，并将结果存储在变量sum中，打印出sum的值

	// 13. 使用defer关键字在函数结束时打印出"Goodbye, Go!"
	defer fmt.Println("Goodbye, Go!")

	// 14. 使用goroutine并发执行一个函数printMessage，该函数输出"Hello, Goroutine!"
	go printMessage()

	time.Sleep(time.Second)
}

func printMessage() {
	fmt.Println("Hello, Goroutine!")
}