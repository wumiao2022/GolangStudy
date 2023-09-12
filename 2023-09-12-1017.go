package main

import (
	"fmt"
)

func main() {
	// 1. 声明一个变量a为整数类型，赋值为10
	var a int = 10
	
	// 2. 声明一个变量b为浮点数类型，赋值为3.14
	var b float64 = 3.14
	
	// 3. 声明一个变量c为字符串类型，赋值为"Hello, World!"
	var c string = "Hello, World!"
	
	// 4. 声明一个变量d为布尔类型，赋值为true
	var d bool = true
	
	// 5. 声明一个变量e为切片类型，包含元素1, 2, 3
	var e []int = []int{1, 2, 3}
	
	// 6. 输出变量a的值
	fmt.Println(a)
	
	// 7. 输出变量b的值
	fmt.Println(b)
	
	// 8. 输出变量c的值
	fmt.Println(c)
	
	// 9. 输出变量d的值
	fmt.Println(d)
	
	// 10. 输出变量e的值
	fmt.Println(e)
}