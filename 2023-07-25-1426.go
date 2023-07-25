package main

import "fmt"

func main() {
	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 将两个整数相加并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 使用循环打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 4. 定义一个结构体Person，并创建一个Person实例并打印其字段值
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person)

	// 5. 使用递归计算斐波那契数列的第n个数，并打印结果
	n := 10
	fibonacci := calcFibonacci(n)
	fmt.Println("Fibonacci(", n, ") =", fibonacci)
}

func calcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calcFibonacci(n-1) + calcFibonacci(n-2)
}