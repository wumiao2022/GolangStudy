package main

import "fmt"

func main() {
    // 1. Hello, World!
    fmt.Println("Hello, World!")

    // 2. Variables
    var name string = "John"
    age := 23
    fmt.Println("My name is", name, "and I am", age, "years old.")

    // 3. Constants
    const pi float64 = 3.14159
    fmt.Println("The value of pi is", pi)

    // 4. Arrays
    numbers := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", numbers)

    // 5. Slices
    letters := []string{"a", "b", "c", "d", "e"}
    fmt.Println("Slice:", letters)

    // 6. Loops
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }

    // 7. Functions
    result := add(3, 5)
    fmt.Println("Result:", result)

    // 8. Structs
    type Person struct {
        name string
        age  int
    }
    person := Person{name: "Alice", age: 30}
    fmt.Println("Person:", person)

    // 9. Pointers
    x := 10
    fmt.Println("Value of x before function call:", x)
    changeValue(&x)
    fmt.Println("Value of x after function call:", x)

    // 10. Interfaces
    animals := []Animal{Cat{}, Dog{}}
    for _, animal := range animals {
        animal.Speak()
    }
}

func add(a, b int) int {
    return a + b
}

func changeValue(x *int) {
    *x = 20
}

type Animal interface {
    Speak()
}

type Cat struct{}

func (c Cat) Speak() {
    fmt.Println("Meow!")
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}