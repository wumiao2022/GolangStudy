package main

import "fmt"

func main() {
	// 1. 使用For循环输出1到10的整数

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 使用For循环计算1到100的累加和

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 3. 使用For循环打印出九九乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 4. 使用For循环判断一个数是否为素数

	num := 17
	isPrime := true

	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}

	// 5. 使用For循环找出一个字符串中的所有元音字母

	str := "hello world"
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	for _, char := range str {
		for _, vowel := range vowels {
			if char == vowel {
				fmt.Printf("%c ", char)
				break
			}
		}
	}

	fmt.Println()
}