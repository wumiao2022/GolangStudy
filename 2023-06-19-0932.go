package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World!")

	// 练习2：输出当前时间
	fmt.Println("Current time is:", time.Now())

	// 练习3：计算两数之和
	num1 := 10
	num2 := 20
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, num1+num2)

	// 练习4：实现一个简单的倒计时器
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Boom!")
}