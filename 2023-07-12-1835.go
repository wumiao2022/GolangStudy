package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1：将字符串全部转换为大写
	str := "hello, world!"
	fmt.Println(strings.ToUpper(str))

	// 练习2：截取字符串的前5个字符
	fmt.Println(str[:5])

	// 练习3：将字符串中的空格替换为横杠
	str2 := "hello world"
	fmt.Println(strings.ReplaceAll(str2, " ", "-"))

	// 练习4：统计字符串中出现的字母个数（不区分大小写）
	str3 := "Hello, World!"
	count := 0
	for _, s := range strings.ToLower(str3) {
		if s >= 'a' && s <= 'z' {
			count++
		}
	}
	fmt.Println(count)
}