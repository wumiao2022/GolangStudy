package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 声明一个字符串变量，打印其长度
	str := "Go programming"
	strLen := len(str)
	fmt.Println("String Length:", strLen)

	// 4. 使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 使用条件语句判断一个数是否为偶数
	num := 7
	if num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
