package main

import "fmt"

func main() {
	// 1. 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 2. 声明一个切片，并使用for循环打印出切片中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}
	
	// 3. 定义一个结构体，并初始化它的字段值
	type Person struct {
		Name string
		Age  int
	}
	
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person)
	
	// 4. 声明一个函数，并在main函数中调用它
	myFunction()
}

func myFunction() {
	fmt.Println("This is my function.")
}
