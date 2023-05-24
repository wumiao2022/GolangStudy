package main

import "fmt"

func main() {
	// 1. 输出Hello World
	fmt.Println("Hello World")

	// 2. 定义变量并输出
	var name string = "Golang"
	fmt.Println("My favorite programming language is", name)

	// 3. 使用for循环输出0-9
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	// 4. 使用if语句判断一个数是否为偶数，是则输出Even，否则输出Odd
	num := 8
	if num%2 == 0 {
		fmt.Println("\nEven")
	} else {
		fmt.Println("\nOdd")
	}

	// 5. 定义一个数组并输出其中的每一个元素
	array := [4]string{"apple", "banana", "cherry", "dragon fruit"}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}

	// 6. 使用range关键字输出数组中的每一个元素
	for index, value := range array {
		fmt.Printf("\narray[%d] = %s", index, value)
	}

	// 7. 定义一个切片并输出其中的每一个元素
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}

	// 8. 使用append函数向切片中添加新元素
	slice = append(slice, 6)
	fmt.Println("\nNew slice after appending 6:", slice)

	// 9. 定义一个map并输出其中的每一个键值对
	m := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"cherry": "red",
	}
	for key, value := range m {
		fmt.Printf("\n%s is %s", key, value)
	}

	// 10. 使用make函数创建一个空的切片，并使用for循环向其中添加元素
	newSlice := make([]int, 0)
	for i := 0; i < 5; i++ {
		newSlice = append(newSlice, i)
	}
	fmt.Println("\nNew slice:", newSlice)
}