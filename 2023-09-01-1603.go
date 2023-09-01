package main

import "fmt"

func main() {
    // 1. 将两个整数进行相加
    add(5, 3)

    // 2. 获取一个整数的绝对值
    abs(-10)

    // 3. 判断一个年份是否是闰年
    isLeapYear(2024)

    // 4. 将一个整数数组进行排序
    array := []int{5, 2, 8, 1, 9}
    sortArray(array)

    // 5. 计算并打印斐波那契数列的前n项
    fibonacci(10)
}

func add(a, b int) {
    sum := a + b
    fmt.Println("Sum:", sum)
}

func abs(num int) {
    if num < 0 {
        num = -num
    }
    fmt.Println("Absolute value:", num)
}

func isLeapYear(year int) {
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year.")
    } else {
        fmt.Println(year, "is not a leap year.")
    }
}

func sortArray(array []int) {
    for i := 0; i < len(array)-1; i++ {
        for j := i + 1; j < len(array); j++ {
            if array[i] > array[j] {
                array[i], array[j] = array[j], array[i]
            }
        }
    }
    fmt.Println("Sorted array:", array)
}

func fibonacci(n int) {
    a, b := 0, 1
    fmt.Print("Fibonacci series:", a, " ", b)
    for i := 2; i < n; i++ {
        sum := a + b
        fmt.Print(" ", sum)
        a, b = b, sum
    }
    fmt.Println()
}