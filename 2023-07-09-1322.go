package main

import "fmt"

func main() {
     // 练习1：打印1到10的数字
     for i := 1; i <= 10; i++ {
          fmt.Println(i)
     }

     // 练习2：计算1到100的和
     sum := 0
     for i := 1; i <= 100; i++ {
          sum += i
     }
     fmt.Println(sum)

     // 练习3：判断一个数是否为偶数
     number := 7
     if number%2 == 0 {
          fmt.Println(number, "是偶数")
     } else {
          fmt.Println(number, "是奇数")
     }

     // 练习4：将一个字符串反转
     str := "Hello, World!"
     reversedStr := ""
     for i := len(str) - 1; i >= 0; i-- {
          reversedStr += string(str[i])
     }
     fmt.Println(reversedStr)
}