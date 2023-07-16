package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组保存一组字符串，并使用循环打印每个字符串
	strings := [4]string{"Go", "is", "awesome", "!"}
	for _, str := range strings {
		fmt.Println(str)
	}
}
