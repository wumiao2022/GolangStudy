package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 变量声明和使用
	var num int
	num = 10
	fmt.Println(num)

	// 3. if条件判断
	if num > 5 {
		fmt.Println("num大于5")
	} else if num == 5 {
		fmt.Println("num等于5")
	} else {
		fmt.Println("num小于5")
	}

	// 4. for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 5. 数组和循环遍历
	arr := [5]string{"apple", "banana", "cherry", "orange", "grape"}
	for _, fruit := range arr {
		fmt.Println(fruit)
	}

	// 6. 函数和返回值
	result := add(5, 3)
	fmt.Println(result)

	// 7. 结构体和方法
	person := Person{name: "Alice", age: 25}
	person.introduce()

	// 8. 切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[1:3])

	// 9. 接口和多态
	var printer Printer
	printer = ConsolePrinter{}
	printer.Print("Hello, Golang!")

	// 10. 并发编程
	ch := make(chan string)
	go sendMessage(ch)
	message := <-ch
	fmt.Println(message)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}

func (p Person) introduce() {
	fmt.Println("My name is", p.name, "and I am", p.age, "years old.")
}

type Printer interface {
	Print(message string)
}

type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(message string) {
	fmt.Println(message)
}

func sendMessage(ch chan<- string) {
	ch <- "Hello, Golang!"
}