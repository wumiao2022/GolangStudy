package main

import "fmt"

func main() {
	// 练习1: 求数组元素之和
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("数组元素之和为：", sum)

	// 练习2: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3: 判断一个数是否为质数
	num := 7
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是质数\n", num)
	} else {
		fmt.Printf("%d不是质数\n", num)
	}
}