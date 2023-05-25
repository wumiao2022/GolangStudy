package main
 
import (
    "fmt"
)
 
func main() {
    // 数组实现冒泡排序
    bubbleSort := func(arr []int) []int {
        for i := 0; i < len(arr); i++ {
            for j := i + 1; j < len(arr); j++ {
                if arr[i] > arr[j] {
                    arr[i], arr[j] = arr[j], arr[i]
                }
            }
        }
        return arr
    }
    arr1 := []int{6, 2, 9, 1, 8, 5, 7, 3, 4}
    fmt.Println(bubbleSort(arr1))
 
    // 使用递归实现斐波那契数列
    fibonacci := func(n int) int {
        if n <= 1 {
            return n
        }
        return fibonacci(n-1) + fibonacci(n-2)
    }
    fmt.Println(fibonacci(7))
 
    // 使用闭包实现累加器
    accumulator := func() func() int {
        sum := 0
        return func() int {
            sum++
            return sum
        }
    }
    a := accumulator()
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
}