package main

import "fmt"

func main() {
	// 1. 声明一个字符串变量，并赋值为 "Hello, Go"
	var str string = "Hello, Go"
	fmt.Println(str)

	// 2. 声明一个整数变量，并赋值为 10
	var num int = 10
	fmt.Println(num)

	// 3. 声明一个布尔变量，并赋值为 true
	var isTrue bool = true
	fmt.Println(isTrue)

	// 4. 声明一个浮点数变量，并赋值为 3.14
	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	// 5. 声明一个切片，并初始化包含有 3 个元素的整数切片
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// 6. 声明一个函数，接收两个整数参数并返回它们的和
	sum := func(a, b int) int {
		return a + b
	}

	result := sum(5, 3)
	fmt.Println(result)

	// 7. 声明一个结构体，包含姓名和年龄字段，并创建一个结构体实例
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "John",
		Age:  25,
	}
	fmt.Println(person)

	// 8. 使用 for 循环输出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 9. 使用 if-else 分支结构判断一个数是否为偶数
	num = 7
	if num % 2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 10. 使用 switch-case 分支结构根据学生成绩进行评级
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
