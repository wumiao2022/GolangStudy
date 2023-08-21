package main

import "fmt"

func main() {
    // 练习1：打印输出Hello World
    fmt.Println("Hello World")

    // 练习2：定义一个整数变量，并打印输出其值
    var num int = 10
    fmt.Println(num)

    // 练习3：定义一个字符串变量，并打印输出其值
    var str string = "Golang"
    fmt.Println(str)

    // 练习4：定义一个浮点数变量，并打印输出其值
    var floatNum float64 = 3.14
    fmt.Println(floatNum)

    // 练习5：定义一个布尔变量，并打印输出其值
    var flag bool = true
    fmt.Println(flag)

    // 练习6：定义一个数组，并打印输出其中元素的值
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)

    // 练习7：定义一个切片，并打印输出其中元素的值
    var slice []int = []int{4, 5, 6}
    fmt.Println(slice)

    // 练习8：定义一个map，并打印输出其中键值对的值
    var m map[string]int = map[string]int{"A": 1, "B": 2, "C": 3}
    fmt.Println(m)

    // 练习9：定义一个结构体，并打印输出其中字段的值
    type Person struct {
        Name string
        Age  int
    }
    p := Person{Name: "Alice", Age: 20}
    fmt.Println(p)

    // 练习10：定义一个函数，并调用该函数
    func printHello() {
        fmt.Println("Hello")
    }
    printHello()
}