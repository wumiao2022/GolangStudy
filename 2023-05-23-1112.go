package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1到100的数字，但是对于3的倍数打印“Fizz”替代数字，对于5的倍数打印“Buzz”替代数字，对于既是3的倍数又是5的倍数的数字打印“FizzBuzz”替代数字
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 练习2：计算斐波那契数列的第n个数字，斐波那契数列从0和1开始，后面每一项数字都是前面两项数字之和，例如0, 1, 1, 2, 3, 5, 8, 13, 21
	fmt.Println(fibonacci(7))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}