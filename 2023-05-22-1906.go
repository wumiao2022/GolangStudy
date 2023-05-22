package main

import (
    "fmt"
)

func main() {
    // 1. 输出Hello World!
    fmt.Println("Hello World!")

    // 2. 计算两个整数的和
    a, b := 10, 20
    sum := a + b
    fmt.Printf("sum=%d\n", sum)

    // 3. 构造一个整型slice，输出其中的每个元素
    nums := []int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }

    // 4. 定义一个结构体Person，包含姓名和年龄两个字段，构造一个Person实例并输出
    type Person struct {
        name string
        age  int
    }
    person := Person{name: "张三", age: 20}
    fmt.Printf("姓名：%s, 年龄：%d\n", person.name, person.age)

    // 5. 定义一个函数，输入两个整数，返回它们的较大值
    getMax := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    fmt.Println(getMax(10, 20))

    // 6. 使用defer打印当前时间
    defer fmt.Println(time.Now())

    // 7. 新建一个缓冲区为1的channel，开启一个goroutine向其中发送数据，然后从中读取数据并输出
    ch := make(chan int, 1)
    go func() {
        ch <- 10
    }()
    fmt.Println(<-ch)

    // 8. 使用Go语言实现快速排序算法
    quickSort := func(nums []int) {
        if len(nums) <= 1 {
            return
        }
        pivotIndex := len(nums) / 2
        pivot := nums[pivotIndex]
        left, right := []int{}, []int{}
        nums = append(nums[:pivotIndex], nums[pivotIndex+1:]...)
        for _, num := range nums {
            if num < pivot {
                left = append(left, num)
            } else {
                right = append(right, num)
            }
        }
        quickSort(left)
        quickSort(right)
        copy(nums[:len(left)], left)
        nums[len(left)] = pivot
        copy(nums[len(left)+1:], right)
    }
    nums = []int{5, 2, 6, 1, 3, 9}
    quickSort(nums)
    fmt.Println(nums)
}