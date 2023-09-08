package main

import "fmt"

func main() {
	// 练习1：打印1到100之间的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2：计算1到100之间所有奇数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习3：打印斐波那契数列的前20项
	a, b := 0, 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 2; i < 20; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}

	// 练习4：实现一个函数，计算两个整数的最大公约数
	fmt.Println(gcd(36, 48))
	fmt.Println(gcd(52, 65))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}