package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：将两个数相加并打印结果
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：将字符串反转并打印结果
	str := "Hello, Go!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)

	// 练习4：打印出从1到100的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习5：计算斐波那契数列的前20个数并打印结果
	fibonacci := []int{0, 1}
	for i := 2; i < 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci sequence:", fibonacci)
}