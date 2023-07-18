package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：打印当前时间
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// 练习4：编写一个函数，判断一个数是否为偶数
	num := 42
	isEven := func(n int) bool {
		return n%2 == 0
	}
	fmt.Println(num, "is even:", isEven(num))

	// 练习5：使用for循环输出1到100之间的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习6：使用if-else语句判断一个数是正数、负数还是零
	num2 := -5
	if num2 > 0 {
		fmt.Println(num2, "is positive")
	} else if num2 < 0 {
		fmt.Println(num2, "is negative")
	} else {
		fmt.Println(num2, "is zero")
	}
}
