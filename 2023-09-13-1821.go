package main

import "fmt"

func main() {
	// 1. 使用for循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 使用if语句判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 3. 创建一个包含5个元素的整数数组，并打印出数组元素的和
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("数组元素的和为:", sum)

	// 4. 定义一个结构体表示人的信息，并创建一个人的实例
	type Person struct {
		Name    string
		Age     int
		Country string
	}

	person := Person{
		Name:    "Alice",
		Age:     25,
		Country: "USA",
	}
	fmt.Println(person)

	// 5. 创建一个函数，接受一个字符串作为参数，并将其进行倒序输出
	reverseString := func(str string) {
		for i := len(str) - 1; i >= 0; i-- {
			fmt.Print(string(str[i]))
		}
		fmt.Println()
	}

	reverseString("Hello")
}