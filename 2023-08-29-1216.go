package main

import "fmt"

func main() {
    // 1. 声明并赋值一个整型变量num
    num := 10

    // 2. 声明一个字符串变量name，并将其初始化为空字符串
    var name string = ""

    // 3. 声明一个浮点型变量pi，并将其初始化为3.14
    var pi float64 = 3.14

    // 4. 声明一个布尔型变量isTrue，并将其初始化为true
    var isTrue bool = true

    // 5. 声明一个切片变量fruits，并使用字面量初始化该切片包含三个字符串元素："apple"、"banana"和"orange"
    fruits := []string{"apple", "banana", "orange"}

    // 6. 打印上述变量的值
    fmt.Println(num)
    fmt.Println(name)
    fmt.Println(pi)
    fmt.Println(isTrue)
    fmt.Println(fruits)
}