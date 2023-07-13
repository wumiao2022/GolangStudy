package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3: 判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	// 练习4: 使用for循环打印0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 使用切片实现栈的基本功能（入栈、出栈、打印栈顶元素）
	stack := make([]int, 0)
	stack = append(stack, 1)      // 入栈
	stack = append(stack, 2)      // 入栈
	top := stack[len(stack)-1]    // 获取栈顶元素
	stack = stack[:len(stack)-1]  // 出栈
	fmt.Println("Top element:", top)

	// 练习6: 使用map实现一个简单的学生管理系统（添加学生、查询学生、删除学生）
	students := make(map[string]int)
	students["Alice"] = 20
	students["Bob"] = 22
	fmt.Println("Age of Alice:", students["Alice"])      // 查询学生
	delete(students, "Bob")                              // 删除学生
	fmt.Println("Number of students:", len(students))    // 查询学生数量
}