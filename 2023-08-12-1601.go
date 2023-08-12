package main

import "fmt"

func main() {
	// 练习1：打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 练习2：打印从0到20的所有偶数
	for i := 0; i <= 20; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 练习3：计算1到100的总和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：打印斐波那契数列的前10个数
	n := 10
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println("Fibonacci:", fibonacci)

	// 练习5：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime?", isPrime)
}