package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World!
	fmt.Println("Hello World!")

	// 练习2：计算并打印1+2的结果
	result := 1 + 2
	fmt.Println("1 + 2 =", result)

	// 练习3：使用for循环计算并打印1+2+3+...+10的结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1+2+3+...+10 =", sum)

	// 练习4：使用if语句判断一个数是否为偶数，如果是，打印"偶数"，否则打印"奇数"
	num := 5
	if num % 2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}