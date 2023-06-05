package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 计算1~100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 3. 判断一个数是偶数还是奇数
    num := 7
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 4. 查找一个切片中的最大值和最小值
    numbers := []int{2, 4, 6, 1, 3, 5}
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
    fmt.Printf("最大值：%d，最小值：%d\n", max, min)

    // 5. 读取一个文件的内容并打印出来
    // 假设文件路径为：test.txt
    // 可以使用 ioutil.ReadFile 方法读取文件内容
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("读取文件出错：", err)
    } else {
        fmt.Println(string(data))
    }
}