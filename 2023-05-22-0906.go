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
    fmt.Println("1~100的和为:", sum)

    // 3. 判断一个数是否为奇数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "是偶数")
    } else {
        fmt.Println(num, "是奇数")
    }

    // 4. 数组翻转
    arr := []int{1, 2, 3, 4, 5}
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    fmt.Println("翻转后的数组为:", arr)

    // 5. 字符串反转
    str := "Hello, World!"
    strLen := len(str)
    for i := 0; i < strLen/2; i++ {
        strBytes := []byte(str) // 字符串的底层是byte数组
        strBytes[i], strBytes[strLen-i-1] = strBytes[strLen-i-1], strBytes[i]
        str = string(strBytes) // 将byte数组转换为字符串
    }
    fmt.Println("反转后的字符串为:", str)
}