package main

import "fmt"

func main() {
    // 1. 输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 输出 1~100 的所有偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 3. 打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, j*i)
        }
        fmt.Println()
    }

    // 4. 数组的遍历
    arr := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 5. 切片的遍历
    slice := []string{"apple", "banana", "orange"}
    for _, v := range slice {
        fmt.Println(v)
    }

    // 6. Map的遍历
    m := map[string]int{"apple": 1, "banana": 2, "orange": 3}
    for k, v := range m {
        fmt.Printf("%s:%d\n", k, v)
    }

    // 7. 方法的调用
    p := Person{name: "Tom", age: 18}
    p.SayHello()
}

// 定义一个Person结构体
type Person struct {
    name string
    age  int
}

// 定义Person的SayHello方法
func (p Person) SayHello() {
    fmt.Printf("My name is %s, I am %d years old\n", p.name, p.age)
}