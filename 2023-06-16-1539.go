package main

import "fmt"

func main() {
	// 练习1：定义一个数组，输出数组中的所有元素
	arr := [5]int{1, 2, 3, 4, 5}

	for _, value := range arr {
		fmt.Println(value)
	}

	// 练习2：定义一个切片，向其中添加元素，并输出切片中的元素
	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5)

	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习3：定义一个map，添加元素，并输出map中的元素
	m := make(map[string]string)

	m["name"] = "Tom"
	m["age"] = "18"

	for key, value := range m {
		fmt.Println(key, value)
	}
}