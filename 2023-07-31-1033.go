package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整型变量num并赋值为10，然后打印出变量num的值
	var num int = 10
	fmt.Println(num)

	// 练习3：声明一个包含5个元素的整型数组arr，初始化为[1, 2, 3, 4, 5]，然后打印出数组的长度和每个元素的值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习4：声明一个切片slice，并通过循环给它赋值为[1, 2, 3, 4, 5]，然后打印出切片的长度和每个元素的值
	slice := make([]int, 5)
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}
	fmt.Println(len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// 练习5：声明一个字典dict，并通过键值对的方式为它赋值为{"name": "John", "age": 25}，然后打印出字典的键和对应的值
	dict := map[string]string{
		"name": "John",
		"age":  "25",
	}
	for key, value := range dict {
		fmt.Println(key, value)
	}
}