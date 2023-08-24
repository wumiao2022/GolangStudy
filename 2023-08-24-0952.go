package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习一：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习二：判断字符串是否以特定的前缀开头
	str := "Golang is awesome!"
	prefix := "Golang"
	result := strings.HasPrefix(str, prefix)
	fmt.Println(result) // 输出true

	// 练习三：将字符串中的字母转换为大写
	str2 := "golang"
	str2 = strings.ToUpper(str2)
	fmt.Println(str2) // 输出GOLANG

	// 练习四：使用for循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习五：使用switch语句判断数字的奇偶性
	num := 7
	switch num % 2 {
	case 0:
		fmt.Println("偶数")
	case 1:
		fmt.Println("奇数")
	}

	// 练习六：使用defer关键字延迟执行函数
	defer fmt.Println("这句话最后输出")
	fmt.Println("这句话先输出")

	// 练习七：使用panic和recover处理异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了一个异常：", err)
		}
	}()
	panic("宕机啦！")
}