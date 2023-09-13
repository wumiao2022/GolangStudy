package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello, World!")

    // 练习2：定义一个整型变量并打印出来
    var num int = 42
    fmt.Println(num)

    // 练习3：定义一个字符串变量并打印出来
    var str string = "Golang"
    fmt.Println(str)

    // 练习4：打印1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：定义一个切片并打印出来
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 练习6：定义一个结构体并打印出来
    type Person struct {
        Name string
        Age  int
    }
    person := Person{Name: "Alice", Age: 25}
    fmt.Println(person)
}