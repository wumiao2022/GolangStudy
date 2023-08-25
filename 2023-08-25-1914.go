package main

import "fmt"

func main() {
    // 练习1：输出Hello World
    fmt.Println("Hello World")

    // 练习2：声明一个整型变量并初始化为10，输出该变量的值
    var num int = 10
    fmt.Println(num)

    // 练习3：声明一个字符串变量并初始化为"Go语言编程"，输出该变量的值
    var str string = "Go语言编程"
    fmt.Println(str)

    // 练习4：声明一个布尔型变量并初始化为true，输出该变量的值
    var flag bool = true
    fmt.Println(flag)

    // 练习5：声明一个整型数组并初始化为{1, 2, 3, 4, 5}，输出数组的长度和第三个元素的值
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(len(arr))
    fmt.Println(arr[2])
}