package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整数变量，并打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，并打印该变量的值
	var str string = "Hello, Golang!"
	fmt.Println(str)

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片，并使用for循环遍历打印切片中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}
}