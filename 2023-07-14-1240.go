package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：变量赋值和打印
	var a int = 10
	var b float64 = 3.14
	var c string = "Golang"
	fmt.Println(a, b, c)

	// 练习3：计算圆的面积和周长
	const PI float64 = 3.14159
	var radius float64 = 5.0
	area := PI * radius * radius
	perimeter := 2 * PI * radius
	fmt.Println("圆的面积：", area)
	fmt.Println("圆的周长：", perimeter)

	// 练习4：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习5：使用for循环打印1到10的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}