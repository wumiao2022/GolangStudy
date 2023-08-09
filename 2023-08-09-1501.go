package main

import "fmt"

func main() {
    // 练习1：创建一个切片，包含10个整数，并打印出切片的长度和容量
    slice := make([]int, 10)
    fmt.Println("切片长度:", len(slice))
    fmt.Println("切片容量:", cap(slice))

    // 练习2：创建一个映射，包含3个学生姓名和对应的成绩，然后遍历映射并打印出学生的姓名和成绩
    scores := map[string]int{
        "Alice":  95,
        "Bob":    80,
        "Charlie": 90,
    }
    for name, score := range scores {
        fmt.Println("学生姓名:", name)
        fmt.Println("学生成绩:", score)
    }

    // 练习3：创建一个结构体类型Person，包含name和age字段，然后创建一个Person类型的变量并初始化它的字段，最后打印出该变量的值
    type Person struct {
        name string
        age  int
    }
    p := Person{name: "Tom", age: 20}
    fmt.Println("人名:", p.name)
    fmt.Println("人年龄:", p.age)
}