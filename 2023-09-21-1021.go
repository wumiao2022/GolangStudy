package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用递归计算阶乘
	fmt.Println("Factorial of 5:", factorial(5))

	// 练习5：使用闭包实现简单的计数器
	counter := getCounter()
	fmt.Println(counter()) // 打印1
	fmt.Println(counter()) // 打印2
	fmt.Println(counter()) // 打印3
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}