package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	var a, b int = 10, 20
	var sum int = a + b
	fmt.Println("a + b =", sum)

	// 练习3：判断一个数是否为偶数
	var num int = 6
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "不是偶数")
	}

	// 练习4：for循环打印1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：切片的基本使用
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	// 练习6：函数的基本使用
	fmt.Println(add(3, 5))
}

// add函数：计算两个数的和
func add(a, b int) int {
	return a + b
}