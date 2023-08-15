package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 求1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 5. 求两个数的最大值
	a, b := 10, 20
	if a > b {
		fmt.Println("Max:", a)
	} else {
		fmt.Println("Max:", b)
	}
}