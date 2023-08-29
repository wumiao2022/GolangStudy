package main

import "fmt"

func main() {
	// 练习1：打印出1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印出1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习3：计算1到10之间所有整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：将一个字符串反转输出
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
}