package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1+2的结果
	fmt.Println(1 + 2)

	// 练习3：声明一个变量name，并将其赋值为你的名字，然后打印出来
	name := "John"
	fmt.Println(name)

	// 练习4：声明一个变量age，并将其赋值为你的年龄，然后打印出来
	age := 25
	fmt.Println(age)

	// 练习5：声明一个数组numbers，包含1、2、3三个元素，并依次打印出来
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
}