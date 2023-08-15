package main

import "fmt"

func main() {
    // 1. 打印出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 声明一个整数变量num，并赋值为10，打印出num的值
    var num int = 10
    fmt.Println(num)

    // 3. 声明一个字符串变量name，并赋值为"GoLang"，打印出name的值
    var name string = "GoLang"
    fmt.Println(name)

    // 4. 声明一个浮点数变量pi，并赋值为3.14，打印出pi的值
    var pi float64 = 3.14
    fmt.Println(pi)

    // 5. 声明一个布尔型变量isTrue，并赋值为true，打印出isTrue的值
    var isTrue bool = true
    fmt.Println(isTrue)

    // 6. 声明一个整数数组numbers，包含5个元素，分别赋值为1, 2, 3, 4, 5，打印出数组numbers的值
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(numbers)

    // 7. 使用for循环遍历数组numbers，打印出每个元素的值
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }

    // 8. 声明一个切片slice，包含3个元素，分别赋值为"a", "b", "c"，打印出切片slice的值
    var slice []string = []string{"a", "b", "c"}
    fmt.Println(slice)

    // 9. 使用for循环遍历切片slice，打印出每个元素的值
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    // 10. 声明一个map对象person，包含name和age两个字段，分别赋值为"John"和30，打印出map对象person的值
    var person map[string]interface{} = map[string]interface{}{"name": "John", "age": 30}
    fmt.Println(person)

    // 11. 使用if条件语句判断num是否等于10，如果相等则打印"num is equal to 10"，否则打印"num is not equal to 10"
    if num == 10 {
        fmt.Println("num is equal to 10")
    } else {
        fmt.Println("num is not equal to 10")
    }

    // 12. 使用switch条件语句判断name的值，如果是"GoLang"则打印"Golang is awesome!"，否则打印"Other programming language"
    switch name {
    case "GoLang":
        fmt.Println("Golang is awesome!")
    default:
        fmt.Println("Other programming language")
    }
}