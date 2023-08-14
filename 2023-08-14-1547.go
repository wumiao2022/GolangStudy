package main

import "fmt"

// 练习1: 打印1到10的数字
func practice1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 练习2: 计算两个整数的和
func practice2(a, b int) int {
	return a + b
}

// 练习3: 判断一个数是否为偶数
func practice3(num int) bool {
	return num%2 == 0
}

// 练习4: 将一个字符串反转输出
func practice4(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 练习5: 计算斐波那契数列的第n个数
func practice5(n int) int {
	if n <= 1 {
		return n
	}
	return practice5(n-1) + practice5(n-2)
}

// 练习6: 判断一个字符串是否是回文字符串
func practice6(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// 主函数
func main() {
	practice1()
	fmt.Println(practice2(3, 5))
	fmt.Println(practice3(10))
	fmt.Println(practice4("hello"))
	fmt.Println(practice5(6))
	fmt.Println(practice6("level"))
}