package main

import "fmt"

func main() {
    // 多维切片的遍历
    arr := [][]int{{1, 2}, {3, 4}, {5, 6}}
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            fmt.Print(arr[i][j], " ")
        }
        fmt.Println()
    }

    // 使用append向切片中添加元素
    s := []int{1, 2, 3}
    s = append(s, 4, 5, 6)
    fmt.Println(s)

    // 使用range遍历切片
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println(sum)

    // 集合的交集
    set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
    set2 := map[int]bool{3: true, 4: true, 5: true, 6: true, 7: true}
    intersection := []int{}
    for k, _ := range set1 {
        if _, ok := set2[k]; ok {
            intersection = append(intersection, k)
        }
    }
    fmt.Println(intersection)
}