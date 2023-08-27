package main

import "fmt"

func main() {
	// 练习1: 打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3: 计算1到100的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4: 计算斐波那契数列的前20个数字
	fib := []int{0, 1}
	for i := 2; i < 20; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println(fib)

	// 练习5: 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is", num, "a prime number?", isPrime)
}
