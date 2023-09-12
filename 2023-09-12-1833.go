package main

import "fmt"

func main() {
	// 练习1: 定义一个函数，接收一个整数参数，判断该整数是否为偶数，并返回结果
	isEven := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(isEven(5))   // false
	fmt.Println(isEven(10))  // true
	fmt.Println(isEven(-3))  // false

	// 练习2: 定义一个结构体，表示一个人的信息，包含姓名、年龄和性别三个字段
	type person struct {
		name   string
		age    int
		gender string
	}
	p1 := person{name: "John", age: 25, gender: "Male"}
	fmt.Println(p1.name)    // John
	fmt.Println(p1.age)     // 25
	fmt.Println(p1.gender)  // Male

	// 练习3: 定义一个切片，包含5个元素，分别为1到5，使用for循环遍历并打印元素值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习4: 定义一个函数，接收一个字符串参数，返回字符串的长度和第一个字符
	getLengthAndFirstChar := func(str string) (int, string) {
		return len(str), string(str[0])
	}
	length, firstChar := getLengthAndFirstChar("Hello")
	fmt.Println(length)     // 5
	fmt.Println(firstChar)  // H

	// 练习5: 定义一个函数，接收两个整数参数，返回这两个整数的乘积和差
	getProductAndDifference := func(num1, num2 int) (int, int) {
		return num1 * num2, num1 - num2
	}
	product, difference := getProductAndDifference(5, 3)
	fmt.Println(product)     // 15
	fmt.Println(difference)  // 2
}