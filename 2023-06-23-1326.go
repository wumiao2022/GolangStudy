package main

import (
	"fmt"
)

func main() {
	// 练习1：输出10以内的偶数
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2：数组遍历
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习3：判断素数
	num := 7
	flag := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习4：字符串遍历
	str := "Hello, world!"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// 练习5：递归求阶乘
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}