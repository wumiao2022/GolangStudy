package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个变量，赋值为10，并打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3：定义两个整型变量a和b，分别赋值为5和7，然后输出它们的和
	a := 5
	b := 7
	fmt.Println(a + b)

	// 练习4：定义一个字符串变量name，赋值为"John"，然后输出"Hello, John!"
	name := "John"
	fmt.Println("Hello, " + name + "!")

	// 练习5：定义一个切片，包含整数1到5，然后遍历切片并打印每个元素的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

注意：以上只是一个示例代码，请将此代码拷贝到一个新的Go文件中，并在自己的开发环境中运行。这些代码都是基础练习，帮助你熟悉基本的Go语法和概念。