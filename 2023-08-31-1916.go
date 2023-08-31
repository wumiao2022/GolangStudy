package main

import "fmt"

func main() {
	// 实例1：打印Hello World
	fmt.Println("Hello World")

	// 实例2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 实例3：判断一个数是否为素数
	number := 17
	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(number, "is prime?", isPrime)

	// 实例4：计算斐波那契数列
	fibonacci := make([]int, 10)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < 10; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println("Fibonacci:", fibonacci)

	// 实例5：求取数组中的最大值
	numbers := []int{4, 2, 9, 6, 5, 10, 8, 3, 7, 1}
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println("Max:", max)
}