package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：声明并打印一个整数变量
	var num int
	num = 10
	fmt.Println(num)

	// 练习3：声明并打印一个字符串变量
	var name string
	name = "John"
	fmt.Println(name)

	// 练习4：计算两个整数的和
	a := 5
	b := 7
	sum := a + b
	fmt.Println(sum)

	// 练习5：判断一个数是否为偶数
	num2 := 12
	if num2%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习6：循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}