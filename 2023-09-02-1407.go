package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 打印当前时间
	fmt.Println(time.Now())

	// 练习2: 求两个整数的最大值
	num1 := 10
	num2 := 20
	max := num1
	if num2 > max {
		max = num2
	}
	fmt.Println("最大值为：", max)

	// 练习3: 判断一个数是否是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "是一个偶数")
	} else {
		fmt.Println(num, "是一个奇数")
	}

	// 练习4: 打印1~100之间的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5: 使用循环计算1~100之间所有偶数的和
	sum := 0
	for i := 2; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println("1~100之间所有偶数的和为：", sum)
}
