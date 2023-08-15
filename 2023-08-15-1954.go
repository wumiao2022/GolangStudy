package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出当前时间
	fmt.Println("Exercise 1:")
	fmt.Println("Current Time:", time.Now())

	// 练习2：计算数组元素的平均值
	fmt.Println("\nExercise 2:")
	numbers := []float64{7, 9, 2, 4, 6, 8}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	fmt.Printf("Average: %.2f\n", average)

	// 练习3：字符串反转
	fmt.Println("\nExercise 3:")
	str := "Golang"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)

	// 练习4：计算两个整数的最大公约数
	fmt.Println("\nExercise 4:")
	num1, num2 := 42, 56
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	fmt.Println("Greatest Common Divisor:", num1)

	// 练习5：输出斐波那契数列
	fmt.Println("\nExercise 5:")
	n := 10
	fmt.Print("Fibonacci Series:", fibonacciSeries(n))
}

func fibonacciSeries(n int) []int {
	fibSeries := []int{0, 1}
	for i := 2; i < n; i++ {
		fibSeries = append(fibSeries, fibSeries[i-1]+fibSeries[i-2])
	}
	return fibSeries
}