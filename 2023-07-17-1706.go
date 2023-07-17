package main

import "fmt"

func main() {
	// 1. 打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 计算两个数的和，并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 判断一个数是否为偶数，并打印结果
	number := 15
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 4. 使用for循环打印10次"Hello, Golang!"
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, Golang!")
	}

	// 5. 定义一个slice，并向其中添加元素，然后打印长度和内容
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Content of slice:", slice)
}