package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2+...+100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime?", isPrime)

	// 练习4：找出一个数组中的最大值
	arr := []int{4, -2, 10, 6, 9}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max value:", max)

	// 练习5：将字符串反转
	str := "Hello, Go!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed string:", reversed)
}