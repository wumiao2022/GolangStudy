package main

import "fmt"

func main() {
	// 练习1：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到100的和并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：打印Fibonacci数列的前10个数字
	fibonacci := []int{0, 1}
	for i := 2; i < 10; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci Sequence:", fibonacci)
}