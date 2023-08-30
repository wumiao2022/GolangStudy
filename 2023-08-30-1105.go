package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并输出其值
	var num1 int = 20
	fmt.Println(num1)

	// 3. 声明一个字符串变量并输出其值
	var str1 string = "Golang"
	fmt.Println(str1)

	// 4. 使用循环求1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 5. 创建一个切片，并遍历输出切片中的元素
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
}