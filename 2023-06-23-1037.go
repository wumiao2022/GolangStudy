package main

import "fmt"

func main() {
    // 1. 数组遍历
    arr := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 2. 切片遍历
    slice := []int{1, 2, 3, 4, 5}
    for i := range slice {
        fmt.Println(slice[i])
    }

    // 3. 字典遍历
    dict := map[string]string{"name": "John", "age": "25"}
    for k, v := range dict {
        fmt.Printf("%s: %s\n", k, v)
    }

    // 4. 函数调用
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(1, 2))

    // 5. 结构体使用
    type Point struct {
        x, y int
    }
    p := Point{1, 2}
    fmt.Println(p.x, p.y)

    // 6. 接口使用
    type Adder interface {
        add(int, int) int
    }
    type Calculator struct{}
    func (c Calculator) add(a, b int) int {
        return a + b
    }
    var a Adder
    a = Calculator{}
    fmt.Println(a.add(1, 2))
}