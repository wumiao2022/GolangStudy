package main

import "fmt"

func main() {
	// 练习1: 打印数字10到1
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习2: 判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 练习4: 定义一个结构体表示学生，包括姓名、学号、年级等字段，并初始化一个学生对象
type Student struct {
	Name  string
	Number int
	Grade string
}

func main() {
	student := Student{
		Name: "Alice",
		Number: 123,
		Grade: "10",
	}
	fmt.Println(student)
}