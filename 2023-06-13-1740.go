package main

import (
	"fmt"
)

func main() {
	// 练习1：打印自己的姓名
	fmt.Println("My name is John Doe.")

	// 练习2：实现一个函数，计算两个整数的和
	sum := add(3, 4)
	fmt.Println(sum)

	// 练习3：实现一个函数，判断一个数是不是偶数
	isEven := checkEven(5)
	fmt.Println(isEven)

	// 练习4：实现一个函数，将一个字符串反转并返回
	reversedStr := reverse("Hello, World!")
	fmt.Println(reversedStr)

	// 练习5：使用for循环输出1~10的平方值
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

func add(a, b int) int {
	return a + b
}

func checkEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}