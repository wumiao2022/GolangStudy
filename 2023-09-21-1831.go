package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并赋值为10，打印此变量的值
	var num int = 10
	fmt.Println(num)

	// 3. 创建一个字符串变量并赋值为"GoLang"，打印此变量的值
	var str string = "GoLang"
	fmt.Println(str)

	// 4. 创建一个整型数组，包含5个元素，分别赋值为1, 2, 3, 4, 5，打印数组
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 5. 声明一个切片并初始化为包含1, 2, 3三个元素的切片，打印切片
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// 6. 使用for循环打印1到10之间的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 7. 创建一个函数，接收两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))
}