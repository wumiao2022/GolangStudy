package main

import (
	"fmt"
)

func main() {
	// 1. 打印出所有的 "even"（偶数）和 "odd"（奇数），分别使用if语句和switch语句实现
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "even")
		default:
			fmt.Println(i, "odd")
		}
	}

	// 2. 将一个整数数组倒序输出
	nums := []int{1, 2, 3, 4, 5}
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(nums[i])
	}

	// 3. 打印出乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 4. 计算斐波那契数列的前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}

	// 5. 判断一个数是否是质数
	num := 7
	isPrime := true
	if num < 2 {
		isPrime = false
	}
	for i := 2; i <= num/2; i++ {
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

	// 6. 求一个数的阶乘
	factorial := 1
	n := 5
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	// 7. 打印出所有的斐波那契数列项直到小于100
	a, b = 0, 1
	for a < 100 {
		fmt.Println(a)
		a, b = b, a+b
	}
}