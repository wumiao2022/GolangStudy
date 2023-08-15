package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到10的和并打印结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 5
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算斐波那契数列的前10个数并打印结果
	fmt.Print("Fibonacci Series for first 10 numbers: ")
	fibonacci := [10]int{0, 1}
	fmt.Print(fibonacci[0], " ", fibonacci[1])
	for i := 2; i < 10; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		fmt.Print(" ", fibonacci[i])
	}
	fmt.Println("")
}

请从上面的例子中学习，编写更多的练习代码，加强你的Golang编程技能。