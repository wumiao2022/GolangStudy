package main

import "fmt"

func main() {
	// 示例1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 示例2：变量和赋值
	var x int = 10
	var y float64 = 3.14
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	
	// 示例3：常量
	const PI float64 = 3.14159
	const isTrue bool = true
	fmt.Println("PI =", PI)
	fmt.Println("isTrue =", isTrue)
	
	// 示例4：条件语句
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are not an adult yet.")
	}
	
	// 示例5：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("The current number is", i)
	}
	
	// 示例6：函数
	fmt.Println("5 + 3 =", add(5, 3))
	
	// 示例7：数组
	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The first element is", nums[0])
	
	// 示例8：切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("The length of the slice is", len(slice))
	
	// 示例9：结构体
	type Person struct {
		name string
		age  int
	}
	person := Person{name: "Alice", age: 25}
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

func add(a, b int) int {
	return a + b
}