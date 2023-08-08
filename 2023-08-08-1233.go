package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明并初始化一个整型变量x，值为10
	x := 10

	// 3. 声明一个字符串变量name，不进行初始化
	var name string

	// 4. 将x的值加1，并将结果赋值给x
	x += 1

	// 5. 将x的值乘2，并将结果赋值给x
	x *= 2

	// 6. 将字符串"Go"赋值给name变量
	name = "Go"

	// 7. 将x和name的值打印出来
	fmt.Println("x:", x)
	fmt.Println("name:", name)

	// 8. 使用if语句判断x的值是否大于20，如果是则打印"x大于20"，否则打印"x小于等于20"
	if x > 20 {
		fmt.Println("x大于20")
	} else {
		fmt.Println("x小于等于20")
	}

	// 9. 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 10. 使用switch语句根据当前时间判断是早上、中午、下午还是晚上，并打印出相应的问候语
	switch now := time.Now().Hour(); {
	case now >= 6 && now < 12:
		fmt.Println("早上好！")
	case now >= 12 && now < 14:
		fmt.Println("中午好！")
	case now >= 14 && now < 18:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}