package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印从1到100的整数，但是对于3的倍数，打印“Fizz”而不是数字，对于5的倍数，打印“Buzz”，同时对于既是3的倍数又是5的倍数的数字，打印“FizzBuzz”。
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 练习2: 实现冒泡排序算法，从小到大排序
	arr := []int{3, 5, 1, 10, 2, 7, 8, 6, 4, 9}
	fmt.Println("排序前：", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后：", arr)

	// 练习3: 实现一个函数，用于计算斐波那契数列的第n项，n从0开始，可以递归或非递归实现
	fmt.Println("第10项斐波那契数列的值为：", fibonacci(10))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}