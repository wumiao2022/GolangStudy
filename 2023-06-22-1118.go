package main

import "fmt"

func main() {
	// 练习1：打印1到100的所有奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：判断一个数是否是素数
	num := 29
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习4：实现一个简单的计算器
	var num1, num2 int
	var operator string
	fmt.Println("请输入两个数字和一个操作符（例如：1 + 2）：")
	fmt.Scanln(&num1, &operator, &num2)

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
		fmt.Println("无法识别的操作符")
	}
}