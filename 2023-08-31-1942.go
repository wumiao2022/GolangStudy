package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. 翻转字符串
	str := "Hello, world!"
	reversedStr := ""
	for _, s := range str {
		reversedStr = string(s) + reversedStr
	}
	fmt.Println(reversedStr)

	// 2. 检查回文字符串
	str2 := "level"
	isPalindrome := true
	for i := 0; i < len(str2)/2; i++ {
		if str2[i] != str2[len(str2)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(isPalindrome)

	// 3. 判断字符串中的元音字母个数
	str3 := "hello world"
	vowelCount := 0
	for _, s := range str3 {
		if unicode.ToLower(s) == 'a' || unicode.ToLower(s) == 'e' || unicode.ToLower(s) == 'i' ||
			unicode.ToLower(s) == 'o' || unicode.ToLower(s) == 'u' {
			vowelCount++
		}
	}
	fmt.Println(vowelCount)

	// 4. 删除字符串中的特定字符
	str4 := "Hello, world!"
	removingChar := ","
	result := strings.ReplaceAll(str4, removingChar, "")
	fmt.Println(result)

	// 5. 拆分字符串为切片
	str5 := "apple,banana,orange"
	separator := ","
	slice := strings.Split(str5, separator)
	fmt.Println(slice)
}