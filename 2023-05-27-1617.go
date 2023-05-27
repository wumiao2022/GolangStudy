package main

import "fmt"

func main() {
	// 1. 输出 "Hello, world!" 到控制台
	fmt.Println("Hello, world!")

	// 2. 将两个数字相加并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 3. 循环遍历打印整数数组
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 4. 判断一个数字是否是偶数并输出结果
	num := 10
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 5. 定义一个函数并调用它
	sayHello()

}

func sayHello() {
	fmt.Println("Hello!")
}