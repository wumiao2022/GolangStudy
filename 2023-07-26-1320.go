package main

import "fmt"

func main() {
    // 实现一个函数，接收一个字符串并将其反转，然后返回反转后的字符串
    reverseString := func(s string) string {
        r := []rune(s)
        for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
            r[i], r[j] = r[j], r[i]
        }
        return string(r)
    }

    fmt.Println(reverseString("Hello, World!"))
}