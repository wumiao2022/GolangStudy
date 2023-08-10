package main

import (
	"fmt"
	"time"
)

func main() {
	// 示例1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 示例2：输出当前时间
	fmt.Println(time.Now())

	// 示例3：计算两个数的和
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 示例4：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 示例5：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例6：定义并使用一个函数
	printMessage("Welcome to Golang!")

	// 示例7：使用切片进行数据存储和操作
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// 示例8：使用map存储键值对
	person := map[string]string{"name": "Alice", "age": "25"}
	fmt.Println("Person:", person)

	// 示例9：使用goroutine实现并发执行
	go printMessage("Hello from Goroutine!")
	fmt.Println("Hello from Main function!")

	// 示例10：使用通道进行数据传输
	message := make(chan string)
	go sendMessage(message)
	receiveMessage(message)
}

// 函数：打印消息
func printMessage(message string) {
	fmt.Println(message)
}

// 函数：发送消息
func sendMessage(ch chan<- string) {
	ch <- "Hello from Channel!"
}

// 函数：接收消息
func receiveMessage(ch <-chan string) {
	fmt.Println(<-ch)
}
