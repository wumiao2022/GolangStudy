package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 9
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：使用循环打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个Slice并添加元素
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	// 练习6：定义一个Map并遍历打印键值对
	person := map[string]string{
		"name":  "John",
		"age":   "30",
		"city":  "New York",
		"email": "john@example.com",
	}
	for key, value := range person {
		fmt.Println(key+":", value)
	}

	// 练习7：使用defer延迟执行一个函数
	defer fmt.Println("Deferred function")

	// 练习8：定义一个结构体并实例化对象
	type Person struct {
		Name string
		Age  int
	}
	person1 := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(person1)
}
```