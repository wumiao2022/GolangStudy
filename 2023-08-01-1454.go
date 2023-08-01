package main

import "fmt"

func main() {
	// 1. 简单的Hello, World程序
	fmt.Println("Hello, World!")

	// 2. 声明变量和使用
	var a int = 10
	var b float64 = 3.14
	var c string = "Golang"
	fmt.Println(a, b, c)

	// 3. 数组的初始化和遍历
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 4. 切片的初始化和操作
	var slice1 []int
	slice1 = append(slice1, 1)
	slice2 := []int{2, 3, 4, 5}
	slice3 := make([]int, len(slice2))
	copy(slice3, slice2)
	fmt.Println(slice1, slice2, slice3)

	// 5. 循环和条件语句
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 6. 函数的定义和调用
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// 7. 结构体的定义和使用
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person)

	// 8. 接口的定义和实现
	type Animal interface {
		Cry()
	}
	type Dog struct{}
	func (d Dog) Cry() {
		fmt.Println("Woof!")
	}
	var animal Animal
	animal = Dog{}
	animal.Cry()
}

func add(a, b int) int {
	return a + b
}