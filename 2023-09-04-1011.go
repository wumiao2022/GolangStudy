package main

import "fmt"

func main() {
    // 实例1：打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 实例2：将两个数相加并打印结果
    a := 5
    b := 10
    sum := a + b
    fmt.Println("Sum =", sum)
    
    // 实例3：打印1到10之间的所有奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }
    
    // 实例4：定义一个结构体并打印其内容
    type Person struct {
        Name  string
        Age   int
        Email string
    }
    
    person := Person{
        Name:  "John",
        Age:   30,
        Email: "john@example.com",
    }
    
    fmt.Println(person)
    
    // 实例5：从命令行中获取用户输入并打印
    var input string
    fmt.Print("请输入一个字符串: ")
    fmt.Scan(&input)
    fmt.Println("您输入的字符串是:", input)
    
    // 实例6：判断一个数是否为素数
    num := 17
    prime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            prime = false
            break
        }
    }
    if prime {
        fmt.Println(num, "是一个素数")
    } else {
        fmt.Println(num, "不是一个素数")
    }
    
    // 实例7：定义一个切片并使用切片操作取出其中的元素
    fruits := []string{"apple", "banana", "orange", "watermelon"}
    slice := fruits[1:3]
    fmt.Println(slice)
}