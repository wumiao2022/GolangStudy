package main

import (
    "fmt"
)

func main() {
    // 1. 打印1到100的数字中，所有能被3整除但不能被5整除的数字
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 != 0 {
            fmt.Println(i)
        }
    }

    // 2. 求出1到100的所有数字的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 3. 将一个int切片中的所有元素按照从大到小的顺序排序
    nums := []int{20, 15, 28, 5, 10}
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] < nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    fmt.Println(nums)

    // 4. 根据给定的年份和月份，计算这个月有多少天。闰年2月有29天
    year, month := 2021, 2
    days := 31
    if month == 4 || month == 6 || month == 9 || month == 11 {
        days = 30
    } else if month == 2 {
        if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
            days = 29
        } else {
            days = 28
        }
    }
    fmt.Printf("%d年%d月有%d天\n", year, month, days)
    
    // 5. 找到一个字符串中出现次数最多的字符和出现的次数
    str := "abbbcddddd"
    m := make(map[byte]int)
    maxCount, maxChar := 0, byte(0)
    for i := 0; i < len(str); i++ {
        m[str[i]]++
        if m[str[i]] > maxCount {
            maxCount = m[str[i]]
            maxChar = str[i]
        }
    }
    fmt.Printf("出现次数最多的字符是%c，出现了%d次\n", maxChar, maxCount)
}