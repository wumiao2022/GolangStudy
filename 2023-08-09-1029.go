package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 输出当前时间
	fmt.Println(time.Now())

	// 2. 定义一个切片，并打印切片的长度和容量
	slice := make([]int, 5, 10)
	fmt.Printf("Length: %d\n", len(slice))
	fmt.Printf("Capacity: %d\n", cap(slice))

	// 3. 使用for循环计算1到100的累加和，并输出结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)

	// 4. 使用switch语句判断当前月份，并输出不同的提示信息
	month := time.Now().Month()
	switch month {
	case time.January:
		fmt.Println("当前是一月份")
	case time.February:
		fmt.Println("当前是二月份")
	case time.March:
		fmt.Println("当前是三月份")
	default:
		fmt.Println("当前是其他月份")
	}

	// 5. 定义一个结构体Person，并实例化一个对象，并输出对象的属性
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
}