package main

import "fmt"

func main() {
    // 练习1：打印1到100的整数，如果是3的倍数则打印Fizz，5的倍数则打印Buzz，既是3的倍数又是5的倍数则打印FizzBuzz
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }

    // 练习2：使用冒泡排序将一个整数切片从小到大排序，并输出结果
    nums := []int{9, 3, 1, 10, 5, 2, 8, 7}
    for i := 0; i < len(nums)-1; i++ {
        for j := 0; j < len(nums)-1-i; j++ {
            if nums[j] > nums[j+1] {
                temp := nums[j]
                nums[j] = nums[j+1]
                nums[j+1] = temp
            }
        }
    }
    fmt.Println(nums)
}