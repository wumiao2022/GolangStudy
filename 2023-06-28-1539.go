package main

import "fmt"

func main() {
	// 练习1: 输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个整型变量并输出
	var num int
	num = 10
	fmt.Println(num)

	// 练习3: 定义一个字符串变量并输出
	var name string
	name = "Golang"
	fmt.Println(name)

	// 练习4: 创建一个数组并输出
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	// 练习5: 创建一个切片并输出
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习6: 使用for循环输出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习7: 使用if语句判断一个数是否为偶数
	num = 6
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习8: 使用switch语句判断一个数是否为正数、负数或零
	num = -5
	switch {
	case num > 0:
		fmt.Println(num, "是正数")
	case num < 0:
		fmt.Println(num, "是负数")
	default:
		fmt.Println(num, "是零")
	}

	// 练习9: 使用函数求两个数的和并输出
	fmt.Println(add(2, 3))

	// 练习10: 使用递归函数计算n的阶乘并输出
	fmt.Println(factorial(5))
}

func add(a, b int) int {
	return a + b
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}