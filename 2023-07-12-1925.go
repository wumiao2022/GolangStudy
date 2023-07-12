package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印两个数的和
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：计算并打印两个数的差
	diff := num2 - num1
	fmt.Println("Difference:", diff)

	// 练习4：计算并打印两个数的乘积
	product := num1 * num2
	fmt.Println("Product:", product)

	// 练习5：计算并打印两个数的商
	quotient := num2 / num1
	fmt.Println("Quotient:", quotient)

	// 练习6：判断一个数是否为偶数并打印结果
	number := 12
	isEven := number%2 == 0
	fmt.Println("Is number even?", isEven)

	// 练习7：判断一个数是否为质数并打印结果
	primeNumber := 17
	isPrime := true
	for i := 2; i < primeNumber; i++ {
		if primeNumber%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is prime number?", isPrime)
}

注意：以上仅为练习代码示例，没有提供解释。如有任何疑问，请随时提问。