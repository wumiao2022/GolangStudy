package main

import "fmt"

func main() {
	// 练习1：打印出1到100的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2：计算并打印出1到100的所有奇数之和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习3：判断一个数是否是质数并打印结果
	num := 97
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习4：求一个数的阶乘并打印结果
	num = 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}