package main

import "fmt"

func main() {
	// 示例1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 示例2
	var num int = 5
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Unknown")
	}

	// 示例3
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	for _, v := range arr1 {
		fmt.Println(v)
	}

	// 示例4
	arr2 := []int{1, 2, 3, 4, 5}
	for i, v := range arr2 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 示例5
	var a int = 10
	var b *int
	b = &a
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Value of b: %d\n", *b)
}