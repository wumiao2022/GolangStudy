package main

import "fmt"

func main() {
	// 实例1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 实例2：求两个数的和并输出
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3：判断一个数是奇数还是偶数
	num := 25
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 实例4：使用循环输出9x9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 实例5：计算斐波那契数列的前n个数
	n := 10
	fib1 := 0
	fib2 := 1
	fmt.Print("Fibonacci Series:", fib1, ", ", fib2)
	for i := 3; i <= n; i++ {
		fib := fib1 + fib2
		fmt.Print(", ", fib)
		fib1 = fib2
		fib2 = fib
	}
	fmt.Println()
}
```