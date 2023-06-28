package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2: 使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3: 定义一个切片，包含10个整数，然后打印出切片中每个元素的值
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习4: 定义一个包含5个字符串的数组，然后遍历数组，打印出每个字符串
	array := [5]string{"apple", "banana", "orange", "grape", "watermelon"}
	for _, value := range array {
		fmt.Println(value)
	}

	// 练习5: 定义一个map，键为字符串类型，值为整型，初始化时添加5个键值对，然后遍历map，打印出每个键值对
	mymap := map[string]int{"apple": 1, "banana": 2, "orange": 3, "grape": 4, "watermelon": 5}
	for key, value := range mymap {
		fmt.Println(key, value)
	}
}