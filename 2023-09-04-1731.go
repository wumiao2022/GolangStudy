package main

import (
	"fmt"
	"math"
)

func main() {
	// 实现一个函数，接收一个参数 n，返回 n 的平方值
	fmt.Println(square(5))

	// 实现一个函数，接收一个参数 x，返回 x 的立方值
	fmt.Println(cube(3))

	// 实现一个函数，接收两个参数 x 和 y，返回它们的最大公约数
	fmt.Println(gcd(10, 15))

	// 实现一个函数，接收一个参数 n，判断 n 是否为素数
	fmt.Println(isPrime(17))

	// 实现一个函数，接收一个参数 n，打印出 n 行杨辉三角
	printYanghuiTriangle(5)
}

func square(n int) int {
	return n * n
}

func cube(x int) int {
	return int(math.Pow(float64(x), 3))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func printYanghuiTriangle(n int) {
	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}