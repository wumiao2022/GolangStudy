package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算两个数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum: ", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, " is even.")
	} else {
		fmt.Println(num, " is odd.")
	}

	// 练习4：循环打印1到10之间的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组存储5个整数，并打印数组的内容
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ", numbers)

	// 练习6：使用切片存储5个整数，并打印切片的内容
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println("Slice: ", slice)

	// 练习7：定义一个结构体表示一个人的基本信息，并打印结构体变量的值
	type Person struct {
		Name    string
		Age     int
		Country string
	}

	person := Person{Name: "John", Age: 30, Country: "USA"}
	fmt.Println("Person: ", person)
}