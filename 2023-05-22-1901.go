package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成一个[0, 100)的随机整数，判断其是偶数还是奇数
	num := rand.Intn(100)
	if num%2 == 0 {
		fmt.Printf("%d是偶数\n", num)
	} else {
		fmt.Printf("%d是奇数\n", num)
	}

	// 练习2：实现一个简单的计算器，可以进行加、减、乘、除运算
	var op rune
	var num1, num2 float64
	fmt.Println("请输入运算符和两个数字，以空格分隔：")
	fmt.Scanln(&op, &num1, &num2)
	switch op {
	case '+':
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case '-':
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case '*':
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case '/':
		if num2 == 0 {
			fmt.Println("除数不能为0")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("无效的运算符")
	}

	// 练习3：将一个华氏温度转换为摄氏温度
	var fahrenheit float64
	fmt.Println("请输入华氏温度：")
	fmt.Scanln(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f℉ = %.2f℃\n", fahrenheit, celsius)
}