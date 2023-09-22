package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：定义一个整型变量num，将其赋值为100，打印出num的值
	num := 100
	fmt.Println(num)

	// 练习3：定义一个字符串变量name，将其赋值为"John"，打印出name的值
	name := "John"
	fmt.Println(name)

	// 练习4：定义一个布尔型变量isTrue，将其赋值为true，打印出isTrue的值
	isTrue := true
	fmt.Println(isTrue)

	// 练习5：定义一个切片变量numbers，存储整数1到5的连续值，打印出numbers的值
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}