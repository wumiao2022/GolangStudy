package main

import "fmt"

func main() {
    // 练习1：打印"Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：声明一个整数变量x，并将其初始化为42，打印出x的值
    var x int = 42
    fmt.Println(x)

    // 练习3：声明一个字符串变量name，并将其初始化为"John Doe"，打印出name的值
    var name string = "John Doe"
    fmt.Println(name)

    // 练习4：声明一个切片变量numbers，并将其初始化为包含整数1、2、3的切片，打印出numbers的值
    var numbers []int = []int{1, 2, 3}
    fmt.Println(numbers)

    // 练习5：声明一个map变量person，并将其初始化为包含"name"为"John"、"age"为25的map，打印出person的值
    var person map[string]interface{} = map[string]interface{}{
        "name": "John",
        "age":  25,
    }
    fmt.Println(person)
}