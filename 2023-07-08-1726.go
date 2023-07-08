package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个字符串变量，并打印出其内容
	var str string = "Hello, Golang!"
	fmt.Println(str)

	// 练习3：声明一个整数数组，并打印出每个元素的值
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习4：声明一个函数，接收两个整数参数，返回它们的和
	fmt.Println(add(10, 20))

	// 练习5：使用循环打印出 1 到 10 的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func add(a, b int) int {
	return a + b
}