package main

import (
	"fmt"
)

func main() {
	// 1. 反转字符串
	str := "hello world"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)

	// 2. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 3. 判断一个数是否是质数
	num := 23
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}

	// 4. 查找一个数组中的最大值和最小值
	arr := []int{3, 6, 2, 8, 1, 10}
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Printf("max=%d, min=%d\n", max, min)
}