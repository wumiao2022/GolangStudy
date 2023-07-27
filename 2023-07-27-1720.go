package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印输出
	sum := 1 + 2
	fmt.Println(sum)

	// 练习3：定义一个整型变量x并赋值为10，打印输出x的值
	x := 10
	fmt.Println(x)

	// 练习4：定义一个字符串变量name并赋值为"John"，打印输出"Hello, John!"
	name := "John"
	fmt.Println("Hello, " + name + "!")

	// 练习5：定义一个整型切片numbers，包含1~5这5个数字，打印输出切片的长度
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(len(numbers))
	
	// 练习6：使用循环遍历切片numbers并打印输出每个元素的值
	for _, num := range numbers {
		fmt.Println(num)
	}
	
	// 练习7：定义一个函数add，接受两个整型参数a和b，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	
	// 练习8：调用函数add，并将结果打印输出
	result := add(3, 4)
	fmt.Println(result)
}