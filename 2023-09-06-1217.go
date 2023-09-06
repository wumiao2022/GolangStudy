package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

```go
package main

import "fmt"

func main() {
    num1 := 10
    num2 := 20

    sum := add(num1, num2)
    fmt.Println("Sum:", sum)
}

func add(a, b int) int {
    return a + b
}
```

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    num := rand.Intn(100)
    fmt.Println("Random number:", num)
}
```