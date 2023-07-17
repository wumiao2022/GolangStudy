package main 

import "fmt"

func main() {
	// 练习1: 定义一个变量并赋值为10，打印变量的值
	var num int = 10
	fmt.Println(num)

	// 练习2: 定义一个数组并初始化，包含5个元素分别为1, 2, 3, 4, 5
	arr := [5]int{1, 2, 3, 4, 5}

	// 练习3: 使用循环遍历数组，并打印每个元素的值
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习4: 定义一个切片并初始化，包含3个元素分别为"apple", "banana", "cherry"
	slice := []string{"apple", "banana", "cherry"}

	// 练习5: 使用循环遍历切片，并打印每个元素的值
	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习6: 定义一个结构体类型Person，包含name和age两个字段
	type Person struct {
		name string
		age  int
	}

	// 练习7: 创建一个Person类型的变量，并给字段赋值，然后打印变量的值
	person := Person{name: "Tom", age: 25}
	fmt.Println(person)

	// 练习8: 定义一个函数Add，接收两个整数参数，返回它们的和
	fmt.Println(Add(3, 5))
}

func Add(a, b int) int {
	return a + b
}