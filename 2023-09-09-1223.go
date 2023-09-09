package main

import "fmt"
  
func main() {
    // 1. 打印Hello, World!
    fmt.Println("Hello, World!")
  
    // 2. 变量和常量
    var x int = 5
    const y string = "Hello"
  
    // 3. 条件语句
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else {
        fmt.Println("x is less than or equal to 10")
    }
  
    // 4. 循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
  
    // 5. 数组和切片
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)
  
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)
  
    // 6. 函数
    result := add(3, 4)
    fmt.Println(result)
}
  
// 函数定义
func add(a, b int) int {
    return a + b
}