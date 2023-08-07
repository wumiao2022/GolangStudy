package main

import "fmt"

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 定义一个整数变量x，并赋值为10，并打印变量的值
    x := 10
    fmt.Println(x)

    // 3. 定义一个字符串变量name，并赋值为"GoLang"，并打印变量的值
    name := "GoLang"
    fmt.Println(name)

    // 4. 定义一个整数变量a，并赋值为5，定义一个整数变量b，并赋值为7，将两个变量的和打印出来
    a := 5
    b := 7
    fmt.Println(a + b)

    // 5. 定义一个浮点数变量pi，并赋值为3.1415，打印变量的值
    pi := 3.1415
    fmt.Println(pi)

    // 6. 定义一个布尔变量isOk，并赋值为true，打印变量的值
    isOk := true
    fmt.Println(isOk)

    // 7. 定义一个数组arr，包含1、2、3三个整数元素，并打印数组的长度
    arr := [3]int{1, 2, 3}
    fmt.Println(len(arr))

    // 8. 定义一个切片slice，包含1、2、3三个整数元素，并打印切片的长度
    slice := []int{1, 2, 3}
    fmt.Println(len(slice))

    // 9. 定义一个map，包含name、age和city三个键值对，并打印map的长度
    m := map[string]interface{}{
        "name": "Alice",
        "age":  25,
        "city": "New York",
    }
    fmt.Println(len(m))

    // 10. 使用条件语句判断一个整数x是否大于10，如果是则打印"大于10"，否则打印"小于或等于10"
    if x > 10 {
        fmt.Println("大于10")
    } else {
        fmt.Println("小于或等于10")
    }
}