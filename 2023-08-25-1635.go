package main

import "fmt"

func main() {
    // 1. 声明一个整数变量x，并将其初始化为10
    var x int = 10

    // 2. 声明一个浮点数变量y，并将其初始化为3.14
    var y float64 = 3.14

    // 3. 声明一个字符串变量name，并将其初始化为"John"
    var name string = "John"

    // 4. 输出变量x的值
    fmt.Println(x)

    // 5. 输出变量y的值
    fmt.Println(y)

    // 6. 输出变量name的值
    fmt.Println(name)

    // 7. 将变量x的值加上20，并将结果保存在x中
    x = x + 20

    // 8. 将变量y的值乘以2.5，并将结果保存在y中
    y = y * 2.5

    // 9. 将字符串"Hello"与变量name拼接起来，并将结果保存在变量greeting中
    greeting := "Hello " + name

    // 10. 输出更新后的变量x的值
    fmt.Println(x)

    // 11. 输出更新后的变量y的值
    fmt.Println(y)

    // 12. 输出变量greeting的值
    fmt.Println(greeting)
}