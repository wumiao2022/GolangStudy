package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：求两个数的和
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println(sum)

	// 练习4：判断一个数是奇数还是偶数
	num := 12
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：使用函数交换两个变量的值
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
