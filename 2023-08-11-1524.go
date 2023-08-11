package main

import "fmt"

func main() {
	// 练习1: 打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2: 定义一个整型变量x，并赋值为10，打印x的值
	var x int = 10
	fmt.Println(x)

	// 练习3: 定义一个字符串变量name，并赋值为"John"，打印name的值
	var name string = "John"
	fmt.Println(name)

	// 练习4: 声明一个数组a，长度为5，元素类型为整型，赋值为[1, 2, 3, 4, 5]，打印数组a的所有元素
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// 练习5: 声明一个切片b，长度为3，元素类型为字符串，赋值为["apple", "banana", "orange"]，打印切片b的所有元素
	var b []string = []string{"apple", "banana", "orange"}
	fmt.Println(b)
}