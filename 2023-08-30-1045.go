package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量声明与赋值
	age := 25
	name := "John"
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 练习3：数组操作
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[3])

	// 练习4：循环控制
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 练习5：函数定义与调用
	add := func(a, b int) int {
		return a + b
	}
	result := add(2, 3)
	fmt.Println(result)

	// 练习6：条件判断
	score := 85
	if score >= 90 {
		fmt.Println("Excellent")
	} else if score >= 80 {
		fmt.Println("Good")
	} else {
		fmt.Println("Average")
	}
}

// 练习7：结构体定义与使用
type Person struct {
	name string
	age  int
}

func (p Person) introduce() {
	fmt.Printf("My name is %s and I am %d years old\n", p.name, p.age)
}

func main() {
	p := Person{"Mary", 30}
	p.introduce()
}