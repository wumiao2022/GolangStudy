package main

import "fmt"

func main() {
	// 练习1：打印从1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	// 练习3：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习5：计算斐波那契数列
	n := 10
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib)
}