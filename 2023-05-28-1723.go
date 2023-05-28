package main
import "fmt"

func main() {
    // Hello World
    fmt.Println("Hello, World!")

    // 多种数据类型
    var a int = 10
    var b float32 = 3.14
    var c bool = true
    var d string = "go"

    // 条件判断
    if a > 0 {
        fmt.Println("a is greater than 0")
    } else {
        fmt.Println("a is less than or equal to 0")
    }

    // 循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 数组
    var arr [3]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    fmt.Println(arr)

    // 切片
    s := []int{1, 2, 3, 4, 5}
    fmt.Println(s)

    // Map
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Println(m)
}