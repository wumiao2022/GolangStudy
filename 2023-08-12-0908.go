package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：使用数组存储并打印1到5的数字
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 练习4：使用if-else语句判断一个数字是否是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：声明一个结构体Person并打印其中的字段值
	type Person struct {
		name string
		age  int
	}

	person := Person{"John", 25}
	fmt.Println(person.name, person.age)
}