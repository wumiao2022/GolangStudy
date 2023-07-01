package main

import "fmt"

func main() {
	// 数组求和
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 切片反转
	letters := []string{"a", "b", "c", "d", "e"}
	reversed := make([]string, len(letters))
	for i, j := 0, len(letters)-1; i < len(letters); i, j = i+1, j-1 {
		reversed[i] = letters[j]
	}
	fmt.Println("Reversed:", reversed)

	// 遍历map
	person := map[string]string{
		"name":  "John",
		"age":   "30",
		"city":  "New York",
		"email": "john@example.com",
	}
	for key, value := range person {
		fmt.Println(key+":", value)
	}

	// 递归阶乘
	fmt.Println("Factorial of 5:", factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}