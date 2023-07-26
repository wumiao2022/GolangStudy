package main

import "fmt"

func main() {
	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 计算两个数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	// 循环遍历字符串
	str := "Golang"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// 判断一个数是偶数还是奇数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 计算斐波那契数列前n项
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci sequence:", fibonacci)
}