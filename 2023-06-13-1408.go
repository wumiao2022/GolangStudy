package main

import (
	"fmt"
)

func main() {
	// 实现一个简单的计算器，支持加法和减法
	num1 := 10
	num2 := 5
	operator := "+"

	if operator == "+" {
		result := num1 + num2
		fmt.Printf("%v %v %v = %v\n", num1, operator, num2, result)
	} else if operator == "-" {
		result := num1 - num2
		fmt.Printf("%v %v %v = %v\n", num1, operator, num2, result)
	} else {
		fmt.Println("Invalid operator: ", operator)
	}

	// 实现一个简单的猜数字游戏，用户每次输入一个数字，直到猜对为止
	number := 10
	guess := 0

	for guess != number {
		fmt.Println("Guess the number: ")
		fmt.Scan(&guess)

		if guess > number {
			fmt.Println("Too high")
		} else if guess < number {
			fmt.Println("Too low")
		} else {
			fmt.Println("You win!")
		}
	}
}