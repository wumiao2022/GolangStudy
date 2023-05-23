package main

import (
    "fmt"
)

func main() {
    // 练习1：定义一个变量，类型为整数，初始值为10，并打印出来
    var num int = 10
    fmt.Println(num)

    // 练习2：定义一个字符串变量，初始值为Hello world，并打印出来
    var str string = "Hello world"
    fmt.Println(str)

    // 练习3：定义一个数组，长度为5，类型为整数，初始值为1,2,3,4,5，并打印出来
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 练习4：定义一个切片，长度为3，类型为字符串，初始值为hello, world, go，并打印出来
    var slice []string = []string{"hello", "world", "go"}
    fmt.Println(slice)

    // 练习5：定义一个结构体，包含name和age两个字段，创建一个结构体变量并初始化，然后打印出来
    type Person struct {
        name string
        age int
    }
    var person1 Person = Person{name: "Tom", age: 20}
    fmt.Println(person1)
}