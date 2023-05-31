package main

import "fmt"

func main() {

	// 1.
	fmt.Println("Hello, World!")

	// 2.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3.
	n := 10
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 4.
	x := 12345
	reverse := 0
	for x > 0 {
		remainder := x % 10
		reverse = reverse*10 + remainder
		x /= 10
	}
	fmt.Println(reverse)

	// 5.
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}