package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成一个长度为10的随机整数切片
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int, 10)
    for i := 0; i < 10; i++ {
        numbers[i] = rand.Intn(100)
    }
    fmt.Println(numbers)

    // 计算切片中的最大值和最小值
    max := numbers[0]
    min := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    fmt.Println("Max:", max)
    fmt.Println("Min:", min)
}

```

