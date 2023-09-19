package main
 
import "fmt"
 
func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")
 
    // 2. 定义变量并输出
    var x int = 5
    fmt.Println(x)
 
    // 3. 使用循环输出1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
 
    // 4. 使用条件语句判断一组数的大小关系
    a := 5
    b := 10
    c := 7
    if a > b {
        fmt.Println("a大于b")
    } else if b > c {
        fmt.Println("b大于c")
    } else {
        fmt.Println("c最大")
    }
 
    // 5. 定义函数并调用
    result := add(3, 7)
    fmt.Println(result)
}
 
func add(a, b int) int {
    return a + b
}