package main

import "fmt"

func main() {
    // 例子1：打印Hello World
    fmt.Println("Hello World")

    // 例子2：求和函数
    nums := []int{1, 2, 3, 4, 5}
    sum := getSum(nums)
    fmt.Println("Sum:", sum)

    // 例子3：计算斐波那契数列
    fib := fibonacci(10)
    fmt.Println("Fibonacci Sequence:", fib)
}

func getSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func fibonacci(n int) []int {
    seq := make([]int, n)
    seq[0] = 0
    seq[1] = 1
    for i := 2; i < n; i++ {
        seq[i] = seq[i-1] + seq[i-2]
    }
    return seq
}