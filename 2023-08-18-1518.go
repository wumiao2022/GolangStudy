package main

import "fmt"

func main() {
	// 1. 实现一个函数，输入一个正整数 n，打印出 1 到 n 的所有奇数
	printOddNumbers(10)

	// 2. 实现一个函数，计算两个整数的和，并返回结果
	sum := add(5, 3)
	fmt.Println(sum)

	// 3. 实现一个函数，输入一个字符串，返回字符串的长度
	length := getLength("Hello, Golang!")
	fmt.Println(length)

	// 4. 实现一个函数，输入一个整数 n，返回 n 的阶乘
	factorial := getFactorial(5)
	fmt.Println(factorial)

	// 5. 实现一个函数，输入一个字符串和一个字符，统计字符在字符串中出现的次数
	count := countChar("Hello, Golang!", 'o')
	fmt.Println(count)
}

// 打印出 1 到 n 的所有奇数
func printOddNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// 计算两个整数的和，并返回结果
func add(a, b int) int {
	return a + b
}

// 返回字符串的长度
func getLength(str string) int {
	return len(str)
}

// 返回 n 的阶乘
func getFactorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

// 统计字符在字符串中出现的次数
func countChar(str string, char byte) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++
		}
	}
	return count
}