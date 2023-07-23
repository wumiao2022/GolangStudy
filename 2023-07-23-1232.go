package main

import "fmt"

func main() {
	// 练习1: 打印1-10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 计算两个整数的和并输出结果
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数并输出结果
	num := 15
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}