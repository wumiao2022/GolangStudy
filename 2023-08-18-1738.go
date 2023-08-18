package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：声明并初始化一个整型变量x，赋值为10，然后将其打印出来。
	x := 10
	fmt.Println(x)

	// 练习3：声明一个字符串变量name，赋值为"GoLang"，然后将其打印出来。
	name := "GoLang"
	fmt.Println(name)

	// 练习4：声明一个bool类型的变量isTrue，赋值为true，然后将其打印出来。
	isTrue := true
	fmt.Println(isTrue)

	// 练习5：声明并初始化一个浮点型变量num1，赋值为3.14，然后将其打印出来。
	num1 := 3.14
	fmt.Println(num1)

	// 练习6：声明两个整型变量a和b，分别赋值为5和7，然后将它们的和打印出来。
	a := 5
	b := 7
	fmt.Println(a + b)

	// 练习7：声明一个包含5个元素的整型数组arr，然后将它们依次打印出来。
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}