package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义并打印一个整数变量
	num := 10
	fmt.Println(num)

	// 练习3：定义并打印一个字符串变量
	str := "Golang is awesome!"
	fmt.Println(str)

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用时间包显示当前的日期和时间
	currentTime := time.Now()
	fmt.Println(currentTime)
}