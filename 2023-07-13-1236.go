package main

import "fmt"

func main() {
    // 实例1：Hello World
    fmt.Println("Hello, World!")

    // 实例2：变量和常量
    var name string = "Alice"
    age := 18
    const pi float64 = 3.14159

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Pi:", pi)

    // 实例3：控制流程语句
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }

    // 实例4：数组和切片
    nums := [5]int{1, 2, 3, 4, 5}
    fmt.Println(nums)

    numsSlice := nums[2:4]
    fmt.Println(numsSlice)

    // 实例5：函数和方法
    sum := func(a, b int) int {
        return a + b
    }

    result := sum(5, 3)
    fmt.Println("Sum:", result)
}

以上为5个练习实例代码，希望对你的Golang学习有所帮助！