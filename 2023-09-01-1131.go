package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：变量与常量
    var name string = "Golang"
    const age int = 10
    fmt.Printf("name: %s, age: %d\n", name, age)

    // 练习3：基本数据类型
    var num1 int = 10
    var num2 float64 = 3.14
    var isTrue bool = true
    fmt.Printf("num1: %d, num2: %f, isTrue: %t\n", num1, num2, isTrue)

    // 练习4：数组和切片
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    slice := []int{6, 7, 8}
    fmt.Println(arr)
    fmt.Println(slice)

    // 练习5：条件语句和循环
    for i := 1; i <= 5; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        } else {
            fmt.Println("Odd")
        }
    }

    // 练习6：函数
    result := add(2, 3)
    fmt.Println(result)

    // 练习7：结构体
    type Person struct {
        name string
        age int
    }
    p1 := Person{"John", 20}
    fmt.Println(p1)
}

func add(a, b int) int {
    return a + b
}