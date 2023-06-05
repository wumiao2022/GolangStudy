package main

import "fmt"

func main() {
    // 1. 打印Hello World
    fmt.Println("Hello World")

    // 2. 定义变量并打印
    var num1 int = 10
    var num2 float32 = 3.14
    var str string = "Golang"
    fmt.Printf("num1=%d, num2=%.2f, str=%s\n", num1, num2, str)

    // 3. 切片的基本操作
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, 3)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)

    // 4. for循环与continue、break的使用
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        if i > 5 {
            break
        }
        fmt.Println(i)
    }

    // 5. map的基本操作
    map1 := make(map[string]string)
    map1["name"] = "Tom"
    map1["age"] = "18"
    fmt.Println(map1)

    // 6. 判断语句if-else的使用
    score := 70
    if score >= 90 {
        fmt.Println("A")
    } else if score >= 80 {
        fmt.Println("B")
    } else if score >= 70 {
        fmt.Println("C")
    } else {
        fmt.Println("D")
    }

    // 7. 函数的定义与调用
    num3, num4 := 10, 20
    sum := add(num3, num4)
    fmt.Println(sum)
}

func add(num1 int, num2 int) int {
    return num1 + num2
}