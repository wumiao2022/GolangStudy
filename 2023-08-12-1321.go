package main

import "fmt"

func main() {
    // 练习1：使用循环打印出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 练习2：使用循环计算1到10的和，并输出结果
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("Sum of 1 to 10:", sum)
    
    // 练习3：使用条件语句判断一个数字是否为偶数，并输出结果
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
    
    // 练习4：使用数组存储一组学生的成绩，并输出每个学生的成绩
    scores := [5]int{85, 76, 92, 88, 79}
    for i := 0; i < len(scores); i++ {
        fmt.Println("Student", i+1, "score:", scores[i])
    }
    
    // 练习5：使用切片对数组进行切割，提取其中的一部分元素，并输出结果
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slicedNumbers := numbers[3:6]
    fmt.Println("Sliced numbers:", slicedNumbers)
}