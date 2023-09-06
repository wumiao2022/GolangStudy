package main

import "fmt"

func main() {
	// 练习1：打印Hello, world!
	fmt.Println("Hello, world!")

	// 练习2：计算1 + 2的结果并打印
	sum := 1 + 2
	fmt.Println(sum)

	// 练习3：将两个整数相加的函数
	result := add(3, 4)
	fmt.Println(result)

	// 练习4：使用循环计算1到10的累加和并打印
	sum = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5：判断一个数是否是偶数并打印结果
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习6：打印一个字符串的长度
	str := "Hello, Golang!"
	fmt.Println(len(str))

	// 练习7：将字符串转换成大写并打印
	fmt.Println(toUpperCase("golang"))

	// 练习8：循环打印1到10的数字，但不包括5
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
	for i := 6; i <= 10; i++ {
		fmt.Println(i)
	}
}

// add函数：将两个整数相加
func add(a, b int) int {
	return a + b
}

// toUpperCase函数：将字符串转换成大写
func toUpperCase(str string) string {
	return strings.ToUpper(str)
}
