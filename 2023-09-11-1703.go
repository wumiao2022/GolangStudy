package main

import "fmt"

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：声明一个整型变量并赋值为10，打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3：声明一个字符串变量并赋值为"Go语言编程"，打印该变量的值
	str := "Go语言编程"
	fmt.Println(str)

	// 练习4：声明一个整型数组并初始化，数组元素为1, 2, 3, 4, 5，打印该数组的值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习5：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}