package main

import "fmt"

// 找出切片中的最大值和最小值
func findMinMax(nums []int) (int, int) {
    min, max := nums[0], nums[0]

    for _, n := range nums {
        if n < min {
            min = n
        }

        if n > max {
            max = n
        }
    }

    return min, max
}

func main() {
    nums := []int{5, 3, 8, 2, 9, 1, 4, 7, 6}

    min, max := findMinMax(nums)
    fmt.Println("最小值:", min)
    fmt.Println("最大值:", max)
}
