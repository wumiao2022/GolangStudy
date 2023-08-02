package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 声明一个整数，并将它打印出来
	var num int = 10
	fmt.Println(num)

	// 练习3: 声明一个字符串数组，并使用循环打印出每个元素
	var fruits [3]string = [3]string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// 练习4: 声明一个切片，并使用循环打印出每个元素
	var numbers []int = []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}

	// 练习5: 使用条件语句判断一个整数是否为正数，并打印出结果
	var n int = -5
	if n > 0 {
		fmt.Println("正数")
	} else {
		fmt.Println("非正数")
	}
}

// 运行结果：

// Hello, World!
// 10
// apple
// banana
// orange
// 1
// 2
// 3
// 4
// 5
// 非正数