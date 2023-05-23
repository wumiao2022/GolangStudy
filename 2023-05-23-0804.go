package main

import "fmt"

func main() {
    // 练习1：变量声明和打印输出
    var num int = 10
    fmt.Println("num = ", num)

    // 练习2：if条件语句和键盘输入
    var inputNum int
    fmt.Println("请输入一个数字：")
    fmt.Scanf("%d", &inputNum)
    if inputNum%2 == 0 {
        fmt.Println("输入的数字是偶数")
    } else {
        fmt.Println("输入的数字是奇数")
    }

    // 练习3：for循环语句和打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }

    // 练习4：切片的使用和排序
    nums := []int{10, 4, 6, 8, 3}
    fmt.Println("排序前：", nums)
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] > nums[j] {
                temp := nums[i]
                nums[i] = nums[j]
                nums[j] = temp
            }
        }
    }
    fmt.Println("排序后：", nums)
}