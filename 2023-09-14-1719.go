package main

import "fmt"

func main() {
	// 1. 使用for循环输出1到10的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 2. 使用for循环计算1到100的所有偶数的和
	sum := 0
	for i := 2; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sum)

	// 3. 使用递归实现阶乘函数
	fmt.Println("Factorial of 5:", factorial(5))

	// 4. 使用switch语句根据给定的月份输出对应的季节
	month := 3
	switch month {
	case 1, 2, 12:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Autumn")
	default:
		fmt.Println("Unknown month")
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}