package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出当前的日期和时间
	currentTime := time.Now()
	fmt.Println("当前时间:", currentTime.Format("2006-01-02 15:04:05"))

	// 练习2：输出自己的名字和年龄
	name := "John"
	age := 25
	fmt.Println("我的名字是", name, "，年龄是", age)

	// 练习3：交换两个变量的值
	a := 10
	b := 20
	fmt.Println("交换前：a =", a, "，b =", b)
	a, b = b, a
	fmt.Println("交换后：a =", a, "，b =", b)

	// 练习4：计算两个数的和、差、积和商
	num1 := 10
	num2 := 5
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	fmt.Println("两数之和:", sum)
	fmt.Println("两数之差:", diff)
	fmt.Println("两数之积:", product)
	fmt.Println("两数之商:", quotient)

	// 练习5：使用for循环打印1到10之间的所有奇数
	fmt.Println("1到10之间的所有奇数：")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}