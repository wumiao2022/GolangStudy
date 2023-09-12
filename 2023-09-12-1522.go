package main

import "fmt"

// 练习1：打印Hello, World!
func main1() {
	fmt.Println("Hello, World!")
}

// 练习2：定义一个变量，存储你的名字，并打印出来
func main2() {
	name := "John"
	fmt.Println(name)
}

// 练习3：定义一个整型数组，包含1到10这十个数字，并打印出数组中的每个元素
func main3() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

// 练习4：实现一个函数，接收一个整型参数n，返回n的平方值
func square(n int) int {
	return n * n
}

// 练习5：调用square函数，计算并打印出2的平方、3的平方和4的平方
func main5() {
	fmt.Println(square(2))
	fmt.Println(square(3))
	fmt.Println(square(4))
}

// 练习6：实现一个函数，接收两个整型参数a和b，返回a和b中较大的那个数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 练习7：调用max函数，计算并打印出5和10中较大的那个数
func main7() {
	fmt.Println(max(5, 10))
}

// 练习8：编写一个循环，打印出1到10之间的所有奇数
func main8() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 练习9：定义一个结构体Person，包含name和age两个字段，并创建一个Person类型的变量，并初始化其字段
type Person struct {
	name string
	age  int
}

// 练习10：创建一个Person类型的变量，并打印出其name和age字段的值
func main10() {
	person := Person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(person.name)
	fmt.Println(person.age)
}

// 练习11：定义一个接口Animal，包含一个方法Speak，该方法无参数，返回值为字符串类型
type Animal interface {
	Speak() string
}

// 练习12：实现Animal接口的方法Speak，并在该方法中返回动物的叫声
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 练习13：创建一个Dog类型的变量，并调用其Speak方法，将其叫声打印出来
func main13() {
	dog := Dog{}
	fmt.Println(dog.Speak())
}