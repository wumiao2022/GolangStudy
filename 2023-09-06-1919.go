package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. 求一个字符串的长度
	str1 := "Hello, Golang!"
	fmt.Println(len(str1))

	// 2. 判断一个字符串是否包含另一个字符串
	str2 := "Golang"
	fmt.Println(strings.Contains(str1, str2))

	// 3. 将字符串转换为大写
	str3 := "golang is awesome"
	fmt.Println(strings.ToUpper(str3))

	// 4. 拼接两个字符串
	str4 := "Hello"
	str5 := "Golang"
	fmt.Println(str4 + ", " + str5)

	// 5. 取子串
	str6 := "Hello, Golang!"
	fmt.Println(str6[7:])

	// 6. 反转字符串
	str7 := "Golang"
	reversedStr := ""
	for _, char := range str7 {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println(reversedStr)

	// 7. 字符串分割
	str8 := "Go is an open source programming language"
	words := strings.Split(str8, " ")
	fmt.Println(words)

	// 8. 字符串替换
	str9 := "Hello, world!"
	newStr9 := strings.Replace(str9, "world", "Golang", -1)
	fmt.Println(newStr9)

	// 9. 判断字符串是否以特定的前缀开头
	str10 := "Hello, Golang!"
	prefix := "Hello"
	fmt.Println(strings.HasPrefix(str10, prefix))

	// 10. 判断字符串是否以特定的后缀结尾
	str11 := "Hello, Golang!"
	suffix := "Golang!"
	fmt.Println(strings.HasSuffix(str11, suffix))
}