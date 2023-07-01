package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明并初始化整型变量x为10，并打印出来
	x := 10
	fmt.Println(x)

	// 3. 声明并初始化字符串变量name为"John"，并打印出来
	name := "John"
	fmt.Println(name)

	// 4. 声明一个切片slice，并将1到5的整数添加到切片中，然后打印切片中的值
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 5. 声明一个映射map，并添加键值对，然后打印映射中的值
	mymap := make(map[string]int)
	mymap["a"] = 1
	mymap["b"] = 2
	fmt.Println(mymap)
}