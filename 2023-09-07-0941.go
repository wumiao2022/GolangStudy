package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")

	// 1. 输出当前时间（格式：年-月-日 时:分:秒）
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 2. 定义一个数组，包含5个整数，并计算它们的和
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 3. 定义一个函数，判断一个整数是否为偶数，是则返回 true，否则返回 false
	isEven := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println("Is 7 even?", isEven(7))

	// 4. 将字符串 "Hello, Go!" 反转后输出
	str := "Hello, Go!"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println("Reversed string:", reversedStr)

	// 5. 定义一个结构体，表示一个人的信息（包含姓名、年龄、性别），并输出
	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	person := Person{
		Name:   "John",
		Age:    25,
		Gender: "Male",
	}
	fmt.Println("Person:", person)
}