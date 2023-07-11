package main

import "fmt"

func main() {
	// 练习1：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算并打印24的阶乘
	result := 1
	for i := 1; i <= 24; i++ {
		result *= i
	}
	fmt.Println(result)

	// 练习3：判断一个数是否为素数
	number := 17
	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(number, "is a prime number:", isPrime)

	// 练习4：打印斐波那契数列的前20个数字
	fibonacci := []int{0, 1}
	for i := 2; i < 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)
}