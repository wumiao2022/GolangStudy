package main

import "fmt"

func main() {
	// 练习1: 使用for循环打印1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 使用if语句判断一个数是否为偶数，如果是则打印 "是偶数"，否则打印 "不是偶数"
	num := 7
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习3: 使用switch语句根据成绩打印对应的等级，90分及以上为A，80分及以上为B，70分及以上为C，60分及以上为D，60分以下为F
	score := 85
	switch {
	case score >= 90:
		fmt.Println("等级：A")
	case score >= 80:
		fmt.Println("等级：B")
	case score >= 70:
		fmt.Println("等级：C")
	case score >= 60:
		fmt.Println("等级：D")
	default:
		fmt.Println("等级：F")
	}

	// 练习4: 定义一个函数，求两个整数的和，并在main函数中调用该函数并打印结果
	result := add(3, 4)
	fmt.Println("结果：", result)
}

func add(a, b int) int {
	return a + b
}