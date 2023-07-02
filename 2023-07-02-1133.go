package main

import "fmt"

func main() {
    // 1. 输出 Hello World!
    fmt.Println("Hello, World!")

    // 2. 计算两个数字的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

    // 3. 循环输出数字 1 到 10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 4. 判断一个数字是否为偶数
    number := 17
    if number%2 == 0 {
        fmt.Println("The number is even")
    } else {
        fmt.Println("The number is odd")
    }

    // 5. 定义并使用一个结构体
    type Person struct {
        name string
        age  int
    }

    person := Person{"John", 25}
    fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}