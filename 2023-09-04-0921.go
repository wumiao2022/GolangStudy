package main

import (
	"fmt"
)

// 题目：实现一个函数，将给定的字符串逆序输出
func reverseString(str string) string {
	runeStr := []rune(str)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

func main() {
	fmt.Println(reverseString("Hello, world!")) // 输出：!dlrow ,olleH
}
