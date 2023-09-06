package main

import (
    "fmt"
)

func main() {
    // 1. 变量声明与赋值
    var num1 int
    num1 = 10

    // 2. 条件语句
    if num1 > 5 {
        fmt.Println("num1 is greater than 5")
    }

    // 3. 循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 4. 数组声明与遍历
    nums := [5]int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }

    // 5. 函数定义与调用
    greeting := getGreeting()
    fmt.Println(greeting)
}

func getGreeting() string {
    return "Hello, Golang!"
}
```

今天的练习实例包含以下内容：

1. 变量声明与赋值
2. 条件语句
3. 循环语句
4. 数组声明与遍历
5. 函数定义与调用

请按照代码所示练习并运行，加深对这些基本概念的理解和熟练度。明天还会有更多练习等待你挑战，继续加油！