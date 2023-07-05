package main

import "fmt"

func main() {
	// 1. 使用变量存储数字并进行计算
	var num1 int = 10
	var num2 int = 5

	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// 2. 使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 使用条件语句判断成绩等级
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 && score < 90 {
		fmt.Println("Grade: B")
	} else if score >= 70 && score < 80 {
		fmt.Println("Grade: C")
	} else if score >= 60 && score < 70 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. 使用数组存储和遍历多个姓名
	names := [3]string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		fmt.Println(name)
	}

	// 5. 使用结构体定义和使用一个人的信息
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

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Gender:", person.Gender)
}