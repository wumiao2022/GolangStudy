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

	// 练习3：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为偶数
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：找出一个整数数组中的最大值
	arr := []int{3, 1, 5, 2, 4}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max:", max)
}