package main

import "fmt"

func main() {
	// 练习1：声明一个整数变量，将其初始化为10，然后将其加上5，并打印出结果
	var num int = 10
	num += 5
	fmt.Println(num)

	// 练习2：声明一个浮点数变量，将其初始化为3.14，然后将其乘以2，并打印出结果
	var pi float64 = 3.14
	pi *= 2
	fmt.Println(pi)

	// 练习3：声明一个字符串变量，将其初始化为"Hello, Golang!"，然后将其转换为大写字母，并打印出结果
	var str string = "Hello, Golang!"
	str = strings.ToUpper(str)
	fmt.Println(str)

	// 练习4：声明一个布尔型变量，将其初始化为true，然后将其取反，并打印出结果
	var flag bool = true
	flag = !flag
	fmt.Println(flag)

	// 练习5：声明一个整型数组，将其初始化为[1,2,3,4,5]，然后遍历数组并打印出每个元素
	nums := [5]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
}