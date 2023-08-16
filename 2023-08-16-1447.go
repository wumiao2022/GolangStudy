package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello World！
	fmt.Println("Hello World!")
	
	// 练习2: 定义一个整型变量num并赋值为10，打印num的值
	num := 10
	fmt.Println(num)
	
	// 练习3: 定义一个浮点型变量pi并赋值为3.14159，打印pi的值
	pi := 3.14159
	fmt.Println(pi)
	
	// 练习4: 定义一个字符串变量name并赋值为"John Smith"，打印name的值
	name := "John Smith"
	fmt.Println(name)
	
	// 练习5: 定义一个布尔型变量isTrue并赋值为true，打印isTrue的值
	isTrue := true
	fmt.Println(isTrue)
	
	// 练习6: 定义一个整型数组arr并初始化为[1, 2, 3, 4, 5]，打印arr的值
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	
	// 练习7: 定义一个切片slice并初始化为[1.1, 2.2, 3.3, 4.4, 5.5]，打印slice的值
	slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(slice)
	
	// 练习8: 定义一个函数add，接收两个整型参数a和b，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 5))
	
	// 练习9: 定义一个结构体Person，包含name和age两个字段，创建一个Person类型的变量并打印其字段值
	type Person struct {
		name string
		age  int
	}
	person := Person{
		name: "Tom",
		age:  20,
	}
	fmt.Println(person.name, person.age)
	
	// 练习10: 使用for循环打印出1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			fmt.Println(i)
		}
	}
}