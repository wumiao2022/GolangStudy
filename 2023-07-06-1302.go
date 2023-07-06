package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 声明一个整数变量，并输出其值
	var num int = 10
	fmt.Println(num)

	// 练习3: 声明一个字符串变量，并输出其值
	var str string = "Golang"
	fmt.Println(str)

	// 练习4: 创建一个切片，并输出其长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 练习5: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6: 使用if-else语句判断一个数是否为奇数
	num = 9
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习7: 使用switch语句判断一个星期几
	day := 5
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
		fmt.Println("无效的日期")
	}
}