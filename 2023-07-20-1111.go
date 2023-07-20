package main

import "fmt"

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并打印结果
	a := 10
	b := 5
	fmt.Println(a + b)

	// 练习3：定义一个切片，并对其进行遍历打印
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习4：创建一个map并进行添加和删除操作
	person := map[string]string{
		"name":   "Alice",
		"age":    "25",
		"gender": "female",
	}
	person["occupation"] = "engineer"
	delete(person, "age")
	fmt.Println(person)

	// 练习5：编写一个函数，接受一个字符串作为参数并返回其长度
	str := "Hello, Golang!"
	fmt.Println(len(str))
}