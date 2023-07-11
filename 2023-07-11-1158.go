package main

import "fmt"

func main() {
	// 打印Hello World
	fmt.Println("Hello World")

	// 数组的遍历
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 切片的操作
	slice := []int{1, 2, 3, 4, 5}

	// 截取切片
	newSlice := slice[1:3]
	fmt.Println(newSlice)

	// 使用range遍历切片
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// 在切片末尾添加元素
	slice = append(slice, 6)
	fmt.Println(slice)

	// 使用map
	ages := make(map[string]int)
	ages["Alice"] = 20
	ages["Bob"] = 25

	// 获取map中的值
	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// 删除map中的值
	delete(ages, "Bob")

	// 判断map中的值是否存在
	_, exists := ages["Bob"]
	fmt.Println("Bob exists:", exists)
}