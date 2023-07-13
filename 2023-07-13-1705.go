package main

import "fmt"

func main() {
	// 练习1: 打印从1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 打印1到100之间所有的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3: 计算1到100之间所有奇数的总和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("奇数的总和为:", sum)

	// 练习4: 判断一个数是否为质数
	num := 23
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是质数")
	} else {
		fmt.Println(num, "不是质数")
	}

	// 练习5: 找出10以内的斐波那契数列
	first := 0
	second := 1
	fmt.Println(first)
	fmt.Println(second)
	for i := 3; i <= 10; i++ {
		next := first + second
		fmt.Println(next)
		first, second = second, next
	}
}