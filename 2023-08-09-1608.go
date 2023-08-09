package main

import "fmt"

func main() {
    // 练习1：输出Hello World
    fmt.Println("Hello World!")

    // 练习2：计算并输出1+2
    fmt.Println(1 + 2)

    // 练习3：使用for循环输出1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习4：定义一个字符串变量并输出
    var str string = "Hello Golang!"
    fmt.Println(str)

    // 练习5：创建一个整型数组并使用for循环输出数组元素
    nums := [5]int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }

    // 练习6：定义一个函数并调用
    sayHello()

    // 练习7：创建一个结构体，并实例化并输出结构体的成员变量
    person := Person{
        name: "John",
        age:  30,
    }
    fmt.Println(person.name, person.age)
}

// 练习6使用的函数
func sayHello() {
    fmt.Println("Hello!")
}

// 练习7使用的结构体
type Person struct {
    name string
    age  int
}