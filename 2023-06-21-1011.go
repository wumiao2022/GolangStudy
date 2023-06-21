package main

import (
    "fmt"
)

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 3. 定义一个整型变量并输出
    var num int = 10
    fmt.Println(num)

    // 4. 定义一个float32类型的变量并输出
    var pi float32 = 3.1415926
    fmt.Println(pi)

    // 5. 使用if语句判断一个数是否为偶数
    num = 5
    if num % 2 == 0 {
        fmt.Println("这个数是偶数")
    } else {
        fmt.Println("这个数是奇数")
    }

    // 6. 定义一个字符串变量并输出
    var str string = "Hello, Golang!"
    fmt.Println(str)

    // 7. 定义一个数组并输出
    var arr [5]int = [5]int {1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 8. 定义一个切片并输出
    var slice []int = []int {1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 9. 使用for循环计算1到10的和
    var sum int = 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 10. 定义一个结构体并输出
    type Person struct {
        name string
        age int
    }
    var p Person = Person{name: "Tom", age: 18}
    fmt.Println(p)
}