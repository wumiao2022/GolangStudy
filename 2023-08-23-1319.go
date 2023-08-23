package main

import (
	"fmt"
)

func main() {
	// 练习1：打印从1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1到100之间所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}

	// 练习4：找出一个切片中的最大值和最小值
	numbers := []int{4, 2, 9, 7, 5, 1}
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}