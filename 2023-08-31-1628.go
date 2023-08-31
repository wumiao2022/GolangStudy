package main

import "fmt"

func main() {
	// 练习1: 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 使用for循环计算1到100的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3: 使用for循环打印出10到1的倒序数字
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习4: 使用for循环打印出2的乘方，从2的0次方到2的10次方
	for i := 0; i <= 10; i++ {
		fmt.Println(2 << i)
	}

	// 练习5: 使用for循环判断一个数是否是质数(除了1和本身，不能被其他数整除)
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime?", isPrime)
}