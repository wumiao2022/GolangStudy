package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")
    
    // 2. 求和
    a, b := 10, 20
    sum := a + b
    fmt.Printf("%d + %d = %d\n", a, b, sum)
    
    // 3. 遍历切片
    arr := []int{1, 2, 3, 4, 5}
    for i, v := range arr {
        fmt.Printf("arr[%d] = %d\n", i, v)
    }
    
    // 4. 求平均数
    scores := []float64{80.0, 90.0, 85.0, 95.0, 70.0}
    var total float64
    for _, score := range scores {
        total += score
    }
    fmt.Printf("平均成绩为：%.2f\n", total/float64(len(scores)))
    
    // 5. 求全年天数
    isLeapYear := func (year int) bool {
        if year%4 == 0 && year%100 != 0 || year%400 == 0 {
            return true
        }
        return false
    }
    
    days := 0
    for i := 1; i <= 12; i++ {
        switch i {
        case 2:
            if isLeapYear(2020) {
                days += 29
            } else {
                days += 28
            }
        case 4, 6, 9, 11:
            days += 30
        default:
            days += 31
        }
    }
    fmt.Printf("2020年共有%d天\n", days)
    
    // 6. 冒泡排序
    nums := []int{45, 23, 78, 56, 89, 12}
    for i := 0; i < len(nums)-1; i++ {
        for j := 0; j < len(nums)-1-i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
    fmt.Println("冒泡排序结果：", nums)
}