package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	fmt.Println(time.Now())

	// 2. 变量交换
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	// 3. 循环打印数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 4. 判断奇偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 5. 求和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 6. 字符串反转
	str := "Hello, world!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 7. 切片操作
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[1:3])
}