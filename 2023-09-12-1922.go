package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1：判断字符串是否以指定前缀开头
	fmt.Println(strings.HasPrefix("hello world", "hello")) // true
	fmt.Println(strings.HasPrefix("hello world", "world")) // false

	// 练习2：判断字符串是否以指定后缀结尾
	fmt.Println(strings.HasSuffix("hello world", "world")) // true
	fmt.Println(strings.HasSuffix("hello world", "hello")) // false

	// 练习3：统计字符串出现的次数
	fmt.Println(strings.Count("hello world", "l")) // 3

	// 练习4：替换字符串中的部分内容
	fmt.Println(strings.Replace("hello world", "world", "golang", -1)) // hello golang

	// 练习5：判断字符串是否包含指定内容
	fmt.Println(strings.Contains("hello world", "world")) // true
	fmt.Println(strings.Contains("hello world", "golang")) // false

	// 练习6：将字符串以指定分隔符分割为多个子串
	fmt.Println(strings.Split("hello,world,golang", ",")) // [hello world golang]

	// 练习7：将字符串所有字符转为小写
	fmt.Println(strings.ToLower("HELLO world")) // hello world

	// 练习8：将字符串所有字符转为大写
	fmt.Println(strings.ToUpper("hello world")) // HELLO WORLD
}