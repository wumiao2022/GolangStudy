package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个整型变量x，并赋值为10，然后将x输出到控制台
	x := 10
	fmt.Println(x)

	// 练习3：定义一个字符串变量name，并赋值为"John Doe"，然后将name输出到控制台
	name := "John Doe"
	fmt.Println(name)

	// 练习4：定义一个整型数组nums，包含5个元素，分别为1, 2, 3, 4, 5，然后将该数组输出到控制台
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 练习5：定义一个切片slice，包含5个元素，分别为6, 7, 8, 9, 10，然后将该切片输出到控制台
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(slice)
}