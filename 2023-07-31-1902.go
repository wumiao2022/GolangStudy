package main

import "fmt"

func main() {
  // 练习1：打印 Hello, World!
  fmt.Println("Hello, World!")

  // 练习2：变量声明与初始化
  var x int = 5
  var y float64 = 3.14
  var z string = "Golang"
  fmt.Println(x, y, z)

  // 练习3：条件语句
  if x > 3 {
    fmt.Println("x is greater than 3")
  } else {
    fmt.Println("x is not greater than 3")
  }

  // 练习4：循环语句
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  // 练习5：切片操作
  numbers := []int{1, 2, 3, 4, 5}
  fmt.Println(numbers[2:4])

  // 练习6：函数定义与调用
  result := add(3, 4)
  fmt.Println(result)
}

func add(a, b int) int {
  return a + b
}
