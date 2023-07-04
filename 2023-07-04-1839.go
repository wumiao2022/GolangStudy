package main

import "fmt"

func main() {
  // 练习1：输出Hello World
  fmt.Println("Hello World")

  // 练习2：定义一个整型变量并赋值为10，然后输出该变量的值
  num := 10
  fmt.Println(num)

  // 练习3：交换两个整型变量的值，不使用临时变量
  a := 5
  b := 10
  a, b = b, a
  fmt.Println(a, b)

  // 练习4：定义一个字符串变量并初始化，然后输出该变量的值的长度
  str := "Hello, Golang"
  fmt.Println(len(str))

  // 练习5：定义一个整型数组，并使用循环遍历数组并输出每个元素的值
  arr := []int{1, 2, 3, 4, 5}
  for _, num := range arr {
    fmt.Println(num)
  }
}