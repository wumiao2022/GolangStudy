package main

import "fmt"

func main() {
  // 示例1: 打印Hello, World!
  fmt.Println("Hello, World!")

  // 示例2: 使用循环输出数字1到10
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }

  // 示例3: 使用条件语句判断一个数是奇数还是偶数
  num := 7
  if num%2 == 0 {
    fmt.Println("偶数")
  } else {
    fmt.Println("奇数")
  }

  // 示例4: 使用数组存储一组数字，并计算它们的总和
  numbers := []int{1, 2, 3, 4, 5}
  sum := 0
  for _, num := range numbers {
    sum += num
  }
  fmt.Println("总和:", sum)

  // 示例5: 使用函数计算两个数的乘积
  result := multiply(5, 7)
  fmt.Println("乘积:", result)
}

// 函数：计算两个数的乘积
func multiply(a, b int) int {
  return a * b
}