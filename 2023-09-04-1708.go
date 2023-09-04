package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整数变量并输出
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量并输出
	str := "Hello, Golang!"
	fmt.Println(str)

	// 练习4：将两个整数相加并输出结果
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习5：使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6：定义一个切片并输出其中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 练习7：定义一个函数，并在主函数中调用该函数
	sayHello()

	// 练习8：使用条件语句判断一个数是否为偶数，并输出结果
	checkEven(4)
	checkEven(7)
}

func sayHello() {
	fmt.Println("Hello, Golang!")
}

func checkEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}