package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习4：输出1到100之间的偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习5：循环输出Golang五次
    for i := 0; i < 5; i++ {
        fmt.Println("Golang")
    }
}

输出结果：
Hello, World!
1
2
3
4
5
6
7
8
9
10
Sum: 5050
2
4
6
8
10
...
Golang
Golang
Golang
Golang
Golang