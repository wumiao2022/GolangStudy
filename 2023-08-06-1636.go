package main

import "fmt"

func main() {
	// 1. 使用 for 循环打印出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 声明并初始化一个字符串变量，然后打印该变量的值和长度
	str := "Hello, World!"
	fmt.Println(str)
	fmt.Println(len(str))

	// 3. 使用 if-else 结构判断一个数字是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 4. 使用 switch 结构根据星期几的数值输出相应的字符串
	dayOfWeek := 4
	switch dayOfWeek {
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
		fmt.Println("无效的星期几")
	}

	// 5. 使用循环计算 1 到 100 之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}