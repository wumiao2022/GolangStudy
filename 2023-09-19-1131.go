package main

import "fmt"

func main() {
	// 练习1：求两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习2：判断一个数是否为偶数并输出结果
	number := 15
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习3：计算并输出斐波那契数列的前 n 个数字
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci sequence:", fibonacci)
}