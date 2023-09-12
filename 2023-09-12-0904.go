package main

import "fmt"

func main() {
    // 练习1: 打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 计算1 + 2的结果并打印
    fmt.Println(1 + 2)

    // 练习3: 声明一个字符串变量并赋值为"Go语言编程"
    msg := "Go语言编程"
    fmt.Println(msg)

    // 练习4: 声明一个整型变量并赋值为10，再声明一个浮点型变量并赋值为3.14，然后将它们相加并打印结果
    num1 := 10
    num2 := 3.14
    sum := float64(num1) + num2
    fmt.Println(sum)

    // 练习5: 循环打印数字1到5
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // 练习6: 声明一个切片并赋值为[1, 2, 3, 4, 5]，然后遍历切片并打印每个元素
    nums := []int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }

    // 练习7: 声明一个map并赋值为{"name": "Alice", "age": 20}，然后遍历map并打印每个键值对
    info := map[string]interface{}{
        "name": "Alice",
        "age":  20,
    }
    for key, value := range info {
        fmt.Println(key, value)
    }
}