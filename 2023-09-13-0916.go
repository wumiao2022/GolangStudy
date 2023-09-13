package main
import "fmt"

func main() {
    // 练习1: 输出Hello World
    fmt.Println("Hello World!")

    // 练习2: 定义并输出一个整数变量
    var num int = 10
    fmt.Println(num)

    // 练习3: 定义并输出一个字符串变量
    var str string = "Golang"
    fmt.Println(str)

    // 练习4: 定义并输出一个布尔变量
    var flag bool = true
    fmt.Println(flag)

    // 练习5: 使用循环输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习6: 定义一个整数切片，包含1到5的元素，进行遍历输出
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Println(num)
    }

    // 练习7: 定义一个结构体类型，包含姓名和年龄两个字段，创建一个结构体实例并输出
    type Person struct {
        Name string
        Age int
    }
    person := Person{Name: "John", Age: 30}
    fmt.Println(person)
}