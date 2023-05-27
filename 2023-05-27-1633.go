package main

import "fmt"

func main() {
   // 例子1：打印Hello World
   fmt.Println("Hello World")
   
   // 例子2：变量声明与赋值
   var a int
   a = 1
   fmt.Println(a)
   
   // 例子3：for循环
   for i:=1;i<=10;i++ {
      fmt.Println(i)
   }
   
   // 例子4：if条件语句
   b := 2
   if b > 1 {
      fmt.Println("b > 1")
   } else {
      fmt.Println("b <= 1")
   }
   
   // 例子5：switch多条件判断
   c := 3
   switch c {
      case 1:
         fmt.Println("c is 1")
      case 2:
         fmt.Println("c is 2")
      default:
         fmt.Println("c is not 1 or 2")
   }
   
   // 例子6：数组遍历
   d := [3]int{1, 2, 3}
   for _,v := range d {
      fmt.Println(v)
   }
   
   // 例子7：函数定义与调用
   e := add(1, 2)
   fmt.Println(e)
}

func add(a int, b int) int {
   return a + b
}