package main

import "fmt"

// Exercise1 请编写一个函数，实现两个整数相加并返回结果
func Exercise1(a, b int) int {
	return a + b
}

// Exercise2 请编写一个函数，实现判断一个整数是否为偶数，是则返回 true，否则返回 false
func Exercise2(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

// Exercise3 请编写一个函数，实现从1累加到n的和并返回结果
func Exercise3(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Exercise4 请编写一个函数，实现将字符串逆序输出
func Exercise4(str string) string {
	runeStr := []rune(str)
	n := len(runeStr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

// Exercise5 请编写一个函数，实现打印九九乘法表
func Exercise5() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	// 测试调用示例代码
	fmt.Println(Exercise1(5, 7))      // 输出: 12
	fmt.Println(Exercise2(10))        // 输出: true
	fmt.Println(Exercise3(100))       // 输出: 5050
	fmt.Println(Exercise4("Hello"))   // 输出: olleH
	Exercise5()                       // 输出: 打印九九乘法表
}
```