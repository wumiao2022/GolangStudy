package main

import "fmt"

func main() {
   // 1. 打印Hello, world!
   fmt.Println("Hello, world!")

   // 2. 声明一个整数变量x，并将其赋值为10
   var x int = 10
   fmt.Println(x)

   // 3. 声明一个字符串变量，并将其赋值为"Go语言"
   var str string = "Go语言"
   fmt.Println(str)

   // 4. 声明一个浮点数变量，并将其赋值为3.14
   var f float64 = 3.14
   fmt.Println(f)

   // 5. 声明一个布尔型变量，并将其赋值为true
   var b bool = true
   fmt.Println(b)

   // 6. 声明一个整型数组，并初始化为{1, 2, 3, 4, 5}
   arr := [5]int{1, 2, 3, 4, 5}
   fmt.Println(arr)

   // 7. 使用for循环打印1到10的数字
   for i := 1; i <= 10; i++ {
      fmt.Println(i)
   }

   // 8. 声明一个切片，并将其元素初始化为{1, 2, 3, 4, 5}
   slice := []int{1, 2, 3, 4, 5}
   fmt.Println(slice)

   // 9. 使用if语句判断x是否大于等于10，并打印相应的结果
   if x >= 10 {
      fmt.Println("x大于等于10")
   } else {
      fmt.Println("x小于10")
   }

   // 10. 声明一个函数add，接收两个整数参数并返回它们的和
   add := func(a, b int) int {
      return a + b
   }

   // 调用add函数并输出结果
   fmt.Println(add(5, 6))
}
