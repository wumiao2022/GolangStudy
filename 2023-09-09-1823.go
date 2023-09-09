package main

import "fmt"

func main() {
	// 以下是多个练习实例代码

	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 实例4：计算斐波那契数列
	fmt.Println(fibonacci(10))

	// 实例5：使用函数闭包
	add := adder()
	for i := 1; i <= 5; i++ {
		fmt.Println(add(i))
	}
}

// 函数：计算斐波那契数列
func fibonacci(n int) []int {
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

// 函数：返回一个闭包函数
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}