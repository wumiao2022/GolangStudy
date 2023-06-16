package main

import "fmt"

func main() {
	// 练习1：打印100以内的奇数
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习2：打印斐波那契数列前20项
	a, b := 1, 1
	fmt.Println(a)
	for i := 1; i < 20; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}

	// 练习3：判断一个数是否为质数
	num := 11
	isPrime := true
	for i := 2; i < num; i++ {
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

	// 练习4：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和为：", sum)
}