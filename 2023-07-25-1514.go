package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Types
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Example 3: Functions
	result := add(3, 5)
	fmt.Println("The sum is:", result)

	// Example 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Conditional Statements
	number := 6
	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// Example 6: Arrays, Slices, and Maps
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	person := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}
```
可以将以上代码拷贝到一个文件中尝试运行。