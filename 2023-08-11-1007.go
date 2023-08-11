package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // 生成一个0到100之间的随机数并打印出来
    num := rand.Intn(101)
    fmt.Println(num)

    // 使用for循环打印出1到10之间的所有奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }

    // 创建一个包含5个元素的字符串数组，并打印出每个元素的长度
    strArr := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
    for _, str := range strArr {
        fmt.Println(len(str))
    }

    // 创建一个map存储学生姓名和对应的年龄，然后遍历打印出每个学生的姓名和年龄
    studentAges := map[string]int{
        "Alice":  20,
        "Bob":    22,
        "Charlie": 19,
        "David":   21,
        "Eve":     18,
    }
    for name, age := range studentAges {
        fmt.Println(name, age)
    }
}
```
```