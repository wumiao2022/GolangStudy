package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello world!
	fmt.Println("Hello world!")

	// 练习2：使用变量输出“我是程序员，我敲代码”
	var job string = "程序员"
	var sentence string = fmt.Sprintf("我是%s，我敲代码", job)
	fmt.Println(sentence)

	// 练习3：使用if语句判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Printf("%d是偶数\n", num)
	} else {
		fmt.Printf("%d是奇数\n", num)
	}

	// 练习4：使用for循环输出1到100的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	// 练习5：使用switch语句判断一个数的大小（0-99）
	num2 := 78
	switch {
	case num2 < 0:
		fmt.Printf("%d是负数\n", num2)
	case num2 >= 0 && num2 <= 99:
		fmt.Printf("%d在0-99之间\n", num2)
	default:
		fmt.Printf("%d大于等于100\n", num2)
	}
	
	// 练习6：使用函数计算两个数之和
	sum := add(3, 5)
	fmt.Printf("3 + 5 = %d\n", sum)
}

func add(num1, num2 int) int {
	return num1 + num2
}