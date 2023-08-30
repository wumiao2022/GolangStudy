package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求和函数
	sum := add(3, 5)
	fmt.Println("Sum:", sum)

	// 练习3：判断是否为素数
	num := 23
	isPrime := checkPrime(num)
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}

	// 练习4：冒泡排序
	arr := []int{5, 1, 4, 2, 8}
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}

// 求和函数
func add(a, b int) int {
	return a + b
}

// 判断是否为素数
func checkPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 冒泡排序
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}