package main

import "fmt"

func main() {
	// 练习1：打印出1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到10之间的整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：输出斐波那契数列的前10项
	fmt.Println("Fibonacci sequence:")
	fib := []int{0, 1}
	for i := 2; i < 10; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println(fib)

	// 练习4：判断一个数是否为素数
	num := 13
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}