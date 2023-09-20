package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算12 + 34的结果并打印
	result := 12 + 34
	fmt.Println(result)

	// 练习3：根据当前时间输出不同的问候语
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 18 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}

	// 练习4：判断一个数是奇数还是偶数并打印结果
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习5：打印1到10之间的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}