package main

import "fmt"

func main() {
	// 1. 使用Go语言打印“Hello, World!”
	fmt.Println("Hello, World!")

	// 2. 定义一个名为add的函数，接收两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(3, 5)) // 输出：8

	// 3. 使用for循环打印1到10的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 4. 定义一个名为Student的结构体类型，包含name和age字段，创建一个Student变量并初始化
	type Student struct {
		name string
		age  int
	}

	student := Student{"Alice", 18}

	fmt.Println(student.name) // 输出：Alice

	// 5. 创建一个切片，包含字符串"apple", "banana", "cherry"
	fruits := []string{"apple", "banana", "cherry"}

	fmt.Println(fruits[0]) // 输出：apple

	// 6. 使用Golang内置的排序函数对切片进行排序
	fruitsSorted := []string{"apple", "banana", "cherry"}

	sort.Strings(fruitsSorted)

	fmt.Println(fruitsSorted) // 输出：[apple banana cherry]

	// 7. 使用Golang内置的map类型，记录学生的成绩，学生名字作为键，成绩作为值
	scores := map[string]int{
		"Alice":   90,
		"Bob":     80,
		"Charlie": 95,
	}

	fmt.Println(scores["Charlie"]) // 输出：95

	// 8. 使用defer关键字，在函数返回前打印一条消息
	defer fmt.Println("Function executed")

	// 输出：Function executed

	// 9. 使用Golang内置的错误处理机制，自定义一个错误并返回
	err := errors.New("custom error")

	fmt.Println(err.Error()) // 输出：custom error

	// 10. 使用Golang内置的协程特性，创建一个新的协程并打印一条消息
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	// 输出：Hello from goroutine

	// 等待协程执行完毕
	time.Sleep(time.Second)
}
```