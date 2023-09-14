package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 输出10以内的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习3: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4: 判断一个数是否是素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习5: 输出Fibonacci数列前20个数
	n := 20
	first, second := 0, 1
	fmt.Println(first, second)
	for i := 3; i <= n; i++ {
		next := first + second
		fmt.Println(next)
		first = second
		second = next
	}
}