package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断是否为素数
	number := 17
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime && number > 1 {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}

	// 练习4：反转字符串
	str := "hello"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println("Reversed string:", reversedStr)

	// 练习5：统计单词个数
	sentence := "Hello world, how are you doing today?"
	words := strings.Split(sentence, " ")
	fmt.Println("Word count:", len(words))
}

注意：以上的代码只是示例，可能并不是最佳实践。