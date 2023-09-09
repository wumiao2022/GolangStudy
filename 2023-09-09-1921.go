package main

import "fmt"

func main() {
    // 1. 将字符串 "Hello, World!" 赋值给变量 message，并打印该变量的值
    message := "Hello, World!"
    fmt.Println(message)

    // 2. 定义一个整型数组 nums，包含数字 1 到 5
    nums := [5]int{1, 2, 3, 4, 5}

    // 3. 使用 for 循环遍历数组 nums，并打印每个元素的值
    for _, num := range nums {
        fmt.Println(num)
    }

    // 4. 定义一个切片 slice，包含数字 6 到 10
    slice := []int{6, 7, 8, 9, 10}

    // 5. 使用 range 关键字遍历切片 slice，并打印每个元素的值
    for _, num := range slice {
        fmt.Println(num)
    }

    // 6. 定义一个包含键值对的 map，键为字符串类型，值为整型
    m := map[string]int{"a": 1, "b": 2, "c": 3}

    // 7. 使用 for-range 循环遍历 map，并打印每个键值对的值
    for key, value := range m {
        fmt.Println(key, value)
    }
}