package main

import "fmt"

func main() {
    // 1. Hello World
    fmt.Println("Hello, World!")

    // 2. 变量
    var a int = 5
    var b = 10
    c := 15
    fmt.Println(a, b, c)

    // 3. 数组
    var arr1 [5]int
    arr2 := [3]int{1, 2, 3}
    fmt.Println(arr1, arr2)

    // 4. 切片
    slice1 := make([]int, 3, 5)
    slice2 := []int{1, 2, 3, 4, 5}
    fmt.Println(slice1, slice2)

    // 5. for循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 6. if-else语句
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }

    // 7. switch语句
    y := 3
    switch y {
    case 1:
        fmt.Println("y is 1")
    case 2:
        fmt.Println("y is 2")
    default:
        fmt.Println("y is not 1 or 2")
    }

    // 8. 函数
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(3, 5))
}