package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1：将字符串 "Hello World" 逆序输出
	str := "Hello World"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	
	// 练习2：将字符串中的奇数位字符删除
	str2 := "Hello World"
	newStr := ""
	for i := 0; i < len(str2); i++ {
		if i%2 == 0 {
			newStr += string(str2[i])
		}
	}
	fmt.Println(newStr)
	
	// 练习3：将字符串中的所有元音字母替换为 "*"
	str3 := "Hello World"
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, vowel := range vowels {
		str3 = strings.ReplaceAll(str3, vowel, "*")
	}
	fmt.Println(str3)
	
	// 练习4：将字符串中的所有单词倒序输出
	str4 := "Hello World"
	words := strings.Split(str4, " ")
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
	fmt.Println()
}