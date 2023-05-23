package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成10个随机数并打印出来
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	// 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 计算斐波那契数列的前20项并打印出来
	fib := make([]int, 20)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < 20; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib)
	
	// 实现一个简单的排序算法：选择排序
	arr := [...]int{5, 2, 9, 1, 7, 6, 4, 3, 8}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	fmt.Println(arr)
}