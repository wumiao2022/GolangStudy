package main

import "fmt"

func main() {
	// 1. 翻转字符串
	str := "Hello, World!"
	reverseString := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseString += string(str[i])
	}
	fmt.Println(reverseString)

	// 2. 判断素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 3. Fibonacci数列
	num1 := 0
	num2 := 1
	fmt.Print(num1, " ", num2)
	for i := 2; i < 10; i++ {
		nextNum := num1 + num2
		fmt.Print(" ", nextNum)
		num1 = num2
		num2 = nextNum
	}
	fmt.Println()

	// 4. 计算阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	// 5. 生成随机数
	// 使用math/rand包
	// 示例代码如下：
	// rand.Seed(time.Now().UnixNano())
	// randomNum := rand.Intn(100)
	// fmt.Println(randomNum)
}