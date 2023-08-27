package main

import "fmt"

func main() {
    // 实例1：打印Hello World
    fmt.Println("Hello World")

    // 实例2：变量的声明和赋值
    var num int
    num = 10
    fmt.Println(num)

    // 实例3：条件语句
    if num > 5 {
        fmt.Println("Num is greater than 5")
    } else {
        fmt.Println("Num is less than or equal to 5")
    }

    // 实例4：循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 实例5：函数
    func add(a, b int) int {
        return a + b
    }

    result := add(3, 5)
    fmt.Println(result)
}

// 实例6：结构体
type Person struct {
    name string
    age  int
}

func main() {
    person := Person{name: "Alice", age: 25}
    fmt.Println(person.name, person.age)
}