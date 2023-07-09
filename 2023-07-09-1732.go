package main

import (
	"fmt"
)

func main() {
	// 练习1：打印出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并输出结果
	fmt.Println(1 + 2)

	// 练习3：声明一个字符串变量并输出
	var message string = "Hello, Go!"
	fmt.Println(message)

	// 练习4：使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：判断一个数是否为偶数并输出结果
	num := 4
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习6：定义一个整型数组并打印出所有的元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}