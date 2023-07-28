package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印
	fmt.Println(1 + 2)

	// 练习3：声明一个整型变量x并赋值为10，打印x的值
	x := 10
	fmt.Println(x)

	// 练习4：声明一个字符串变量name并赋值为"GoLang"，打印name的值
	name := "GoLang"
	fmt.Println(name)

	// 练习5：声明一个整型数组numbers，并初始化为{1, 2, 3, 4, 5}，打印数组元素的值
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 练习6：使用for循环打印数组numbers的每个元素
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习7：从键盘输入一个整数n，使用if语句判断n的奇偶性，并打印结果
	var n int
	fmt.Print("请输入一个整数n：")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("n是偶数")
	} else {
		fmt.Println("n是奇数")
	}
}