package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个整数的和并打印结果
	a := 6
	b := 9
	sum := a + b
	fmt.Println("Sum:", sum)

	// 3. 定义一个字符串变量，并打印出其长度
	str := "Golang"
	fmt.Println("Length:", len(str))

	// 4. 使用for循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 定义一个数组，并打印出数组的所有元素
	arr := [...]int{1, 2, 3, 4, 5}
	for _, num := range arr {
		fmt.Println(num)
	}
}
