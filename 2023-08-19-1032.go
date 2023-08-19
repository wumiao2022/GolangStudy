package main

import "fmt"

func main() {
	// 练习1
	fmt.Println("Hello, world!")

	// 练习2
	num1 := 10
	num2 := 5
	fmt.Println(num1 + num2)

	// 练习3
	name := "Alice"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习4
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// 练习5
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 练习6
	animals := map[string]string{
		"dog":   "bark",
		"cat":   "meow",
		"horse": "neigh",
	}
	for key, value := range animals {
		fmt.Printf("The sound of a %s is %s.\n", key, value)
	}
}
