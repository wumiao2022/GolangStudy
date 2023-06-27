package main

import "fmt"

func main() {
	// 练习1：输出1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：输出5的倍数，1-50范围内
	for i := 1; i <= 50; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3：输出斐波那契数列前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}

	// 练习4：求1-100的整数和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5：求100以内的质数
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}