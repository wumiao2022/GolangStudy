package main

import "fmt"

func main() {
    // 1. 打印"Hello, world!"
    fmt.Println("Hello, world!")
    
    // 2. 声明并打印变量a的值为10
    var a int = 10
    fmt.Println(a)
    
    // 3. 打印变量b的类型和值
    b := "Golang"
    fmt.Printf("%T %v\n", b, b)
    
    // 4. 多重赋值交换变量c和d的值
    c, d := 1, 2
    c, d = d, c
    fmt.Println(c, d)
    
    // 5. 声明slice e，长度为5，容量为10，打印长度、容量和值
    e := make([]int, 5, 10)
    fmt.Println(len(e), cap(e), e)
    
    // 6. 如果a的值大于10，输出"I'm a good programmer"，否则输出"I need more practice"
    if a > 10 {
        fmt.Println("I'm a good programmer")
    } else {
        fmt.Println("I need more practice")
    }
    
    // 7. 使用for循环打印1-10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 8. 声明函数add，传入两个int类型参数，返回它们的和
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(1, 2))
    
    // 9. 打印15的二进制、八进制和十六进制表示
    num := 15
    fmt.Printf("%b %o %x\n", num, num, num)
    
    // 10. 声明struct person，包含name和age两个字段，并打印出值
    type person struct {
        name string
        age int
    }
    p := person{"Tom", 18}
    fmt.Println(p)
}