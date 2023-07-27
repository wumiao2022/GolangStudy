package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 输出1到100的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3: 将一个字符串反转
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
}