package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：输出数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1到100的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}

	// 练习5：计算斐波那契数列的前n项
	n := 10
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println("Fibonacci sequence:", fib)
}