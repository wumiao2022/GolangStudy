package main

import "fmt"

func main() {
	// 练习一：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习二：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习三：判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 练习四：遍历打印一个切片中的每个元素
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		fmt.Println(n)
	}

	// 练习五：获取字符串的长度
	str := "Hello, Golang!"
	length := len(str)
	fmt.Println("The length of the string is", length)
}