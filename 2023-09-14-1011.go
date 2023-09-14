package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 输出1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习3: 计算100以内所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sum)

	// 练习4: 输出斐波那契数列的前20个数字
	fmt.Print("Fibonacci Series: ")
	fibonacci(20)
}

// 计算斐波那契数列
func fibonacci(n int) {
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		fmt.Print(first, " ")
		temp := first + second
		first = second
		second = temp
	}
}
