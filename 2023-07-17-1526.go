package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：求1到100的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：打印出0到10之间的所有偶数
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习4：打印出10到0的所有数字（递减）
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

	// 练习5：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}
