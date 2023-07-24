package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 实例1：生成一个随机数，判断其是否为偶数，并输出结果
	num := rand.Intn(100)
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 实例2：使用for循环打印出1到10的乘法表
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}

	// 实例3：实现一个简单的计算器，可以进行基本的加减乘除运算
	num1 := 10
	num2 := 5
	operator := "+"

	switch operator {
	case "+":
		fmt.Println(num1, "+", num2, "=", num1+num2)
	case "-":
		fmt.Println(num1, "-", num2, "=", num1-num2)
	case "*":
		fmt.Println(num1, "*", num2, "=", num1*num2)
	case "/":
		fmt.Println(num1, "/", num2, "=", num1/num2)
	default:
		fmt.Println("Invalid operator.")
	}
}