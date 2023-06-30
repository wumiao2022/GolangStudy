package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成10个1到100之间的随机整数，并将它们打印出来
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100) + 1)
	}

	// 练习2：将一个字符串反转，并将结果打印出来
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 练习3：统计一个字符串中每个字符出现的次数，并将结果打印出来
	str2 := "abracadabra"
	charCount := make(map[byte]int)
	for i := 0; i < len(str2); i++ {
		char := str2[i]
		charCount[char]++
	}
	for char, count := range charCount {
		fmt.Printf("%c: %d\n", char, count)
	}
}