package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前的日期和时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 定义一个切片并初始化
	nums := []int{1, 2, 3, 4, 5}
	// 循环遍历切片并输出每个元素
	for _, num := range nums {
		fmt.Println(num)
	}

	// 使用map存储一些键值对
	user := map[string]string{
		"name":  "Tom",
		"email": "tom@example.com",
		"phone": "123456789",
	}
	// 循环遍历map并输出每个键值对
	for k, v := range user {
		fmt.Printf("%s: %s\n", k, v)
	}
}