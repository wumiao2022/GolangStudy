package main 

import (
    "fmt"
)

func main() {
    //1. 输出Hello World!
    fmt.Println("Hello World!")
    
    //2. 定义一个字符串变量并赋值，然后将其打印出来
    str := "Golang is amazing!"
    fmt.Println(str)
    
    //3. 定义一个整型变量并赋值，然后将其打印出来
    num := 7
    fmt.Println(num)
    
    //4. 定义一个浮点型变量并赋值，然后将其打印出来
    flo := 3.14
    fmt.Println(flo)
    
    //5. 定义一个布尔型变量并赋值，然后将其打印出来
    bol := true
    fmt.Println(bol)
    
    //6. 定义一个切片并打印出来
    slc := []int{1,2,3,4,5}
    fmt.Println(slc)
    
    //7. 定义一个map并打印出来
    m := map[string]int{"a":1, "b":2, "c":3}
    fmt.Println(m)
    
    //8. 使用if语句判断一个数是否大于10并打印出结果
    x := 11
    if x > 10 {
        fmt.Println("x is greater than 10.")
    } else {
        fmt.Println("x is not greater than 10.")
    }
    
    //9. 使用for循环打印出1到10的数字
    for i:=1; i<=10; i++ {
        fmt.Println(i)
    }
    
    //10. 定义一个结构体并打印出来
    type person struct {
        name string
        age int
        gender string
    }
    p := person{name:"Tom", age:24, gender:"male"}
    fmt.Println(p)
}