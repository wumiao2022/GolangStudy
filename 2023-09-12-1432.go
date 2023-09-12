package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：声明一个变量，赋予其整数值，然后打印
    var num int = 5
    fmt.Println(num)

    // 练习3：声明一个切片，并添加元素，然后打印
    var slice []int
    slice = append(slice, 1)
    slice = append(slice, 2)
    fmt.Println(slice)

    // 练习4：声明一个map，并添加键值对，然后打印
    var myMap map[string]int
    myMap = make(map[string]int)
    myMap["a"] = 1
    myMap["b"] = 2
    fmt.Println(myMap)
    
    // 练习5：使用for循环输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}