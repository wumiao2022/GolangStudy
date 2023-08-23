package main

import "fmt"

func main() {
	// 1. 将两个整数相加并打印结果
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 2. 将一个字符串拆分成字符并逐个打印
	str := "Hello, World!"
	for _, char := range str {
		fmt.Println(string(char))
	}

	// 3. 判断一个数是否为素数，并打印结果
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is Prime:", isPrime)

	// 4. 将一个整数逆序输出
	num = 12345
	reversedNum := 0
	for num > 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}
	fmt.Println("Reversed Number:", reversedNum)

	// 5. 打印1到100之间的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
