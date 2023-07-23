package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1+2的结果
	fmt.Println(1 + 2)

	// 练习3：定义一个字符串变量，然后打印出来
	str := "Hello, Golang!"
	fmt.Println(str)

	// 练习4：定义一个整型变量和一个浮点型变量，然后将它们相加并打印结果
	num1 := 10
	num2 := 3.14
	sum := float64(num1) + num2
	fmt.Println(sum)

	// 练习5：使用for循环打印出1到10的所有数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6：定义一个切片，并使用for循环打印出切片中的所有元素
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习7：定义一个函数，接受两个参数并返回它们的和
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}