package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1+2的结果
	result := 1 + 2
	fmt.Println(result)

	// 练习3：使用for循环打印1到10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个切片，包含1, 2, 3, 4, 5这五个元素，并打印切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 练习5：定义一个结构体Person，包含姓名和年龄字段，并创建一个Person类型的变量，并打印其值
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person)
}