package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	a := 2
	b := 3
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// 练习3：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习4：使用for循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：计算1到100的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("The sum of numbers from 1 to 100 is", sum)
}