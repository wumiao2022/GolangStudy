package main

import (
    "fmt"
)

func main() {
    // 1. 声明一个整数变量并赋值为42
    var num int = 42
    fmt.Println(num)

    // 2. 声明一个浮点数变量并赋值为3.14
    var fnum float64 = 3.14
    fmt.Println(fnum)

    // 3. 声明一个字符串变量并赋值为"Hello, World!"
    var str string = "Hello, World!"
    fmt.Println(str)

    // 4. 声明一个布尔型变量并赋值为true
    var flag bool = true
    fmt.Println(flag)

    // 5. 声明一个整数切片并赋值为{1, 2, 3}
    var nums []int = []int{1, 2, 3}
    fmt.Println(nums)

    // 6. 声明一个字典并赋值为{"name": "Tom", "age": 18}
    var person map[string]interface{} = map[string]interface{}{
        "name": "Tom",
        "age":  18,
    }
    fmt.Println(person)

    // 7. 声明一个函数并调用它
    hello("Golang")

    // 8. 声明一个结构体并赋值为{name: "Alice", age: 20}
    var student struct {
        name string
        age  int
    }
    student.name = "Alice"
    student.age = 20
    fmt.Println(student)
}

func hello(name string) {
    fmt.Println("Hello, " + name + "!")
}