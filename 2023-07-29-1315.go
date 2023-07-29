package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：变量和常量
	var message string = "Hello, Golang!"
	const pi float64 = 3.14159

	// 实例3：循环和条件语句
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(message)
		}
	}

	// 实例4：数组和切片
	numbers := [3]int{1, 2, 3}
	numbers[0] = 4
	slice := numbers[1:3]

	// 实例5：函数和方法
	sum := add(2, 3)
	myStruct := MyStruct{value: "example"}
	s := myStruct.getValue()

	// 实例6：接口和类型断言
	var myInterface MyInterface = MyStruct{}
	switch myInterface.(type) {
	case MyStruct:
		fmt.Println("MyStruct type")
	default:
		fmt.Println("Unknown type")
	}

	// 实例7：并发
	ch := make(chan int)
	go doSomething(ch)
	result := <-ch

	// 实例8：错误处理
	_, err := doSomethingWithError()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 示例函数：相加两个整数
func add(a, b int) int {
	return a + b
}

// 示例结构体
type MyStruct struct {
	value string
}

// 示例方法：获取结构体的值
func (m MyStruct) getValue() string {
	return m.value
}

// 示例接口
type MyInterface interface {
	getValue() string
}

// 示例并发函数
func doSomething(ch chan<- int) {
	// 执行一些任务
	result := 42

	// 将结果发送到通道
	ch <- result
}

// 示例错误处理函数
func doSomethingWithError() (string, error) {
	// 执行一些任务
	// ...
	return "", fmt.Errorf("something went wrong")
}