package main

import "fmt"

func main() {

	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 定义并打印两个整数相加的结果
	a := 5
	b := 10
	fmt.Println(a + b)

	// 3. 定义一个字符串变量，并打印其长度
	str := "Hello, Golang!"
	fmt.Println(len(str))

	// 4. 定义一个浮点数数组，并打印其中的元素
	floatArray := []float64{1.2, 3.4, 5.6, 7.8}
	fmt.Println(floatArray)

	// 5. 使用for循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. 使用if语句判断一个数是否为正数，并打印判断结果
	num := 10
	if num > 0 {
		fmt.Println("是正数")
	} else {
		fmt.Println("不是正数")
	}

	// 7. 定义一个结构体类型，并创建一个结构体变量
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Tom", 20}
	fmt.Println(person)

	// 8. 使用switch语句根据数字的值打印对应的星期几
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("无效的数字")
	}

	// 9. 使用数组切片实现队列的入队和出队操作
	queue := []int{}
	queue = append(queue, 1)         // 入队
	first := queue[0]               // 取队首元素
	queue = queue[1:]               // 出队
	fmt.Println(first, queue)

	// 10. 使用map存储学生的成绩，并打印学生的姓名和对应的成绩
	scores := map[string]int{
		"Tom": 78,
		"Bob": 90,
		"Alice": 85,
	}
	for name, score := range scores {
		fmt.Println(name, score)
	}
}