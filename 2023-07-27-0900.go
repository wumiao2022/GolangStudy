package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量并打印其值
	var num int
	num = 10
	fmt.Println(num)

	// 练习3：使用循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：声明一个字符串变量并打印其长度
	str := "Golang"
	fmt.Println(len(str))

	// 练习5：使用条件语句判断一个整数是否为正数
	num = -5
	if num > 0 {
		fmt.Println("正数")
	} else {
		fmt.Println("非正数")
	}

	// 练习6：使用数组存储五个学生成绩，并计算平均成绩
	var scores [5]float64
	scores[0] = 90
	scores[1] = 80
	scores[2] = 85
	scores[3] = 95
	scores[4] = 70

	sum := 0.0
	for _, score := range scores {
		sum += score
	}
	average := sum / float64(len(scores))
	fmt.Println("平均成绩:", average)

	// 练习7：使用map存储五个学生的姓名和年龄，并根据姓名查找年龄
	students := map[string]int{
		"Alice":  20,
		"Bob":    18,
		"Charlie": 22,
		"David":  19,
		"Eve":    21,
	}

	name := "David"
	age, ok := students[name]
	if ok {
		fmt.Println(name, "的年龄是", age)
	} else {
		fmt.Println("无法找到", name, "的年龄信息")
	}
}