package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：声明一个整型变量并初始化为10，打印它的值
	var num int = 10
	fmt.Println(num)

	// 练习4：声明一个字符串变量并初始化为空字符串，打印它的值
	var str string = ""
	fmt.Println(str)

	// 练习5：使用if-else语句判断一个数是否为偶数
	num = 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}