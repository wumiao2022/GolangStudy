package main

import "fmt"

// 练习1：打印Hello World
func exercise1() {
	fmt.Println("Hello World")
}

// 练习2：计算两个整数相加的和
func exercise2(a, b int) int {
	return a + b
}

// 练习3：将字符串逆序输出
func exercise3(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 练习4：判断一个数是否是质数
func exercise4(num int) bool {
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

// 练习5：计算斐波那契数列的第n项
func exercise5(n int) int {
	if n <= 1 {
		return n
	}
	return exercise5(n-1) + exercise5(n-2)
}

func main() {
	// 每日练习示例
	exercise1()
	fmt.Println(exercise2(3, 5))
	fmt.Println(exercise3("Hello World"))
	fmt.Println(exercise4(17))
	fmt.Println(exercise5(8))
}
