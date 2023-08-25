package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello World!
	fmt.Println("Hello World!")

	// 练习2：打印 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算 1 到 100 的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为偶数
	num := 12
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：遍历字符串并打印每个字符的ASCII码值
	str := "Hello"
	for _, c := range str {
		fmt.Println(c, ":", string(c))
	}
}