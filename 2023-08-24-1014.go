package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum:", sum)
}
```

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is not greater than y")
	}
}
```

```go
package main

import "fmt"

func main() {
	var radius float64

	fmt.Println("Please enter the radius:")
	fmt.Scanln(&radius)

	area := 3.14 * (radius * radius)
	circumference := 2 * 3.14 * radius

	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Circumference: %.2f\n", circumference)
}
```