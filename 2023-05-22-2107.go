package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1到10的和为：", sum)

	// 练习3：打印日期和时间
	currentTime := time.Now()
	fmt.Println("当前日期和时间为：", currentTime.Format("2006-01-02 15:04:05"))

	// 练习4：判断一个数是否为偶数
	number := 10
	if number%2 == 0 {
		fmt.Println(number, "是偶数")
	} else {
		fmt.Println(number, "是奇数")
	}

	// 练习5：计算一个数的阶乘
	number2 := 5
	factorial := 1
	for i := 1; i <= number2; i++ {
		factorial *= i
	}
	fmt.Println(number2, "的阶乘为：", factorial)
}