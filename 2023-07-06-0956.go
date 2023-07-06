package main

import "fmt"

func main() {
    // 1. 打印Hello, World!
    fmt.Println("Hello, World!")

    // 2. 声明一个整型变量x，并赋值为10
    x := 10

    // 3. 声明一个字符串变量name，并赋值为"John"
    name := "John"

    // 4. 打印x的值和name的值
    fmt.Println("x =", x)
    fmt.Println("name =", name)

    // 5. 声明一个切片numbers，并初始化为包含1、2、3三个元素的切片
    numbers := []int{1, 2, 3}

    // 6. 打印numbers的长度和容量
    fmt.Println("len(numbers) =", len(numbers))
    fmt.Println("cap(numbers) =", cap(numbers))

    // 7. 声明一个map students，键为学号，值为学生姓名
    students := map[int]string{
        1: "Alice",
        2: "Bob",
        3: "Charlie",
    }

    // 8. 打印students中学号为2的学生姓名
    fmt.Println("students[2] =", students[2])

    // 9. 循环打印numbers中的每个元素
    for _, num := range numbers {
        fmt.Println(num)
    }

    // 10. 声明一个函数add，接受两个整型参数，返回它们的和
    add := func(a, b int) int {
        return a + b
    }

    // 11. 调用函数add，并打印结果
    sum := add(5, 3)
    fmt.Println("sum =", sum)
}
