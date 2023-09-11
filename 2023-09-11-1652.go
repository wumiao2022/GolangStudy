package main

import "fmt"

func main() {
    // 练习1：打印 Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：变量值交换
    a := 10
    b := 20
    a, b = b, a
    fmt.Println("a =", a, "b =", b)
    
    // 练习3：判断奇偶数
    num := 15
    if num%2 == 0 {
        fmt.Println(num, "是偶数")
    } else {
        fmt.Println(num, "是奇数")
    }
    
    // 练习4：遍历切片
    names := []string{"Alice", "Bob", "Charlie", "Dave"}
    for i, name := range names {
        fmt.Println(i+1, name)
    }
    
    // 练习5：计算阶乘
    n := 5
    factorial := 1
    for i := 1; i <= n; i++ {
        factorial *= i
    }
    fmt.Println(n, "的阶乘是", factorial)
}