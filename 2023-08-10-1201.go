package main

import "fmt"

func main() {
	// 练习1
	var num1 int = 10
	var num2 int = 20
	fmt.Println(num1 + num2)

	// 练习2
	var str1 string = "Hello"
	var str2 string = "World"
	fmt.Println(str1 + str2)

	// 练习3
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// 练习4
	cities := []string{"Beijing", "Shanghai", "Guangzhou"}
	for _, city := range cities {
		fmt.Println(city)
	}

	// 练习5
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}
}