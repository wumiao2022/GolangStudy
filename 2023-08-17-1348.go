package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 实例2：声明并打印一个整数变量
	var num int = 10
	fmt.Println(num)
	
	// 实例3：使用if语句判断一个数是正数、负数还是零
	num = -5
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
	
	// 实例4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 实例5：声明并打印一个字符串切片
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)
}