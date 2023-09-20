package main

import "fmt"

func main() {
    // 练习1：变量和赋值
    var x int = 7
    fmt.Println(x)

    // 练习2：条件语句
    if x > 5 {
        fmt.Println("x is greater than 5")
    }

    // 练习3：循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 练习4：切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers)

    // 练习5：函数
    multiply(2, 3)

    // 练习6：结构体
    person := Person{"John", 25}
    fmt.Println(person)

    // 练习7：指针
    p := &x
    fmt.Println(*p)
}

func multiply(a, b int) {
    fmt.Println(a * b)
}

type Person struct {
    name string
    age  int
}