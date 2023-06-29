package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100 is:", sum)

	// 练习3：判断一个数是否是质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 练习4：查找字符串中的特定字符
	str := "Golang is awesome"
	char := 'a'
	count := 0
	for _, ch := range str {
		if ch == char {
			count++
		}
	}
	fmt.Printf("Character '%c' appears %d times in the string\n", char, count)
}