package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang每日多练")

	// 1. 定义一个 int 类型的变量 num，赋值为 10
	num := 10
	fmt.Println(num)

	// 2. 定义一个 string 类型的变量 str，赋值为 "Hello, World!"
	str := "Hello, World!"
	fmt.Println(str)

	// 3. 定义一个 float64 类型的变量 price，赋值为 9.99
	price := 9.99
	fmt.Println(price)

	// 4. 创建一个数组 arr，包含 5 个 int 类型的元素，值为 1, 2, 3, 4, 5
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 5. 创建一个切片 slice，包含 3 个 string 类型的元素，值为 "apple", "banana", "cherry"
	slice := []string{"apple", "banana", "cherry"}
	fmt.Println(slice)

	// 6. 创建一个 map 类型的变量 m，包含两个键值对，分别为 "name":"John" 和 "age":25
	m := map[string]interface{}{
		"name": "John",
		"age":  25,
	}
	fmt.Println(m)

	// 7. 使用 for 循环打印出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 8. 使用时间戳获取当前时间，并打印出来
	now := time.Now()
	fmt.Println(now)

	// 9. 使用 if-else 判断 num 是否大于 5，如果是则打印 "num 大于 5"，否则打印 "num 不大于 5"
	if num > 5 {
		fmt.Println("num 大于 5")
	} else {
		fmt.Println("num 不大于 5")
	}

	// 10. 创建一个函数 add，接受两个 int 类型的参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))
}