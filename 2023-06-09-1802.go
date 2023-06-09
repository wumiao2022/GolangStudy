package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println("当前时间：", time.Now())

	// 1. 打印 1~10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 打印 hello world 10 次
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello, World!")
	}

	// 3. 判断一个数字是否是偶数
	num := 20
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "不是偶数")
	}

	// 4. 计算两个数字的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(num1, "+", num2, "=", sum)

	// 5. 反转一个字符串
	str := "hello world"
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
}