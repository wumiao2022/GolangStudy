package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 输出当前时间
	fmt.Println("Current time:", time.Now())

	// 3. 数组操作
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 4. 切片操作
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// 5. for循环
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. if-else语句
	num := 10
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 7. switch语句
	day := time.Now().Weekday()
	switch day {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is a weekday")
	}

	// 8. 函数
	result := sum(2, 3)
	fmt.Println("Sum:", result)

	// 9. 结构体
	person := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}

// 8. 函数定义
func sum(a, b int) int {
	return a + b
}

// 9. 结构体定义
type Person struct {
	Name string
	Age  int
}