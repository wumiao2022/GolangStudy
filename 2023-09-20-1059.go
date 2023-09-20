package main

import "fmt"

func main() {
	// 练习1：打印 Hello World
	fmt.Println("Hello World")

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 9
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：判断一个年份是否为闰年
	year := 2021
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}

	// 练习5：计算斐波那契数列的前10个数
	fmt.Println("Fibonacci series of first 10 numbers:")
	fib1, fib2 := 0, 1
	fmt.Println(fib1)
	fmt.Println(fib2)
	for i := 3; i <= 10; i++ {
		fib := fib1 + fib2
		fmt.Println(fib)
		fib1 = fib2
		fib2 = fib
	}
}