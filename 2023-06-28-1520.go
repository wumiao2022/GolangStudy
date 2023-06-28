package main

import "fmt"

func main() {
    // 示例1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 示例2：计算两个数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

    // 示例3：判断一个数是否为偶数
    num := 15
    if num%2 == 0 {
        fmt.Printf("%d is an even number\n", num)
    } else {
        fmt.Printf("%d is an odd number\n", num)
    }

    // 示例4：循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 示例5：定义一个结构体类型并创建实例
    type Person struct {
        name string
        age  int
    }
    p1 := Person{name: "John", age: 30}
    fmt.Printf("Name: %s, Age: %d\n", p1.name, p1.age)
}
```