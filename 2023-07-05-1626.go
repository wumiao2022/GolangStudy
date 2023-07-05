package main

import "fmt"

func main() {
	// 练习1：打印出Hello World!
	fmt.Println("Hello World!")

	// 练习2：定义一个整数变量，并将其赋值为10，然后打印出来
	var num int
	num = 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，并将其赋值为"Hello, Golang!"，然后打印出来
	var str string
	str = "Hello, Golang!"
	fmt.Println(str)

	// 练习4：定义一个浮点数变量，并将其赋值为3.14，然后打印出来
	var floatNum float64
	floatNum = 3.14
	fmt.Println(floatNum)

	// 练习5：定义一个布尔型变量，并将其赋值为true，然后打印出来
	var isTrue bool
	isTrue = true
	fmt.Println(isTrue)
}

// 练习6：编写一个函数，接受一个整数参数n，打印出从1到n的所有整数
func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

// 练习7：编写一个函数，接受一个整数参数n，计算并返回1到n的所有整数的和
func sumNumbers(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// 练习8：定义一个结构体类型Person，包含name和age两个字段，并打印出一个Person实例的信息
type Person struct {
	name string
	age  int
}

func printPersonInfo(person Person) {
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

// 练习9：定义一个接口类型Animal，包含一个方法Speak()，返回动物的叫声
type Animal interface {
	Speak() string
}

// 练习10：定义一个结构体类型Dog，实现Animal接口的Speak()方法
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 练习11：定义一个结构体类型Cat，实现Animal接口的Speak()方法
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}