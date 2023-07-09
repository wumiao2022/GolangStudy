package main

import (
	"fmt"
)

func main() {
	// 实例1：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 实例2：打印10到1的数字
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 实例3：计算1到100的和并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 实例4：判断一个数是否是奇数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 实例5：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}