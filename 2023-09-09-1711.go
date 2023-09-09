package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 打印当前时间
	fmt.Println(time.Now())

	// 练习2: 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 练习3: 根据用户输入的年龄判断是否成年
	var age int
	fmt.Println("请输入您的年龄:")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("您已经成年!")
	} else {
		fmt.Println("您还未成年!")
	}

	// 练习4: 遍历一个整数切片并打印每个元素的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5: 根据给定的字母判断是否为元音字母，并输出结果
	letter := 'a'
	if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
		fmt.Println("该字母是元音字母")
	} else {
		fmt.Println("该字母不是元音字母")
	}
}