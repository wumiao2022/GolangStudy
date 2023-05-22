package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：定义一个整型变量，并打印出来
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，并打印出来
	str := "Golang"
	fmt.Println(str)

	// 练习4：定义一个浮点型变量，并打印出来
	floatNum := 3.1415
	fmt.Println(floatNum)

	// 练习5：定义一个布尔型变量，并打印出来
	boolValue := true
	fmt.Println(boolValue)

	// 练习6：定义一个数组，包含5个整型元素，并打印出来
	intArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	// 练习7：定义一个切片，包含3个字符串元素，并打印出来
	strSlice := []string{"apple", "banana", "orange"}
	fmt.Println(strSlice)

	// 练习8：定义一个map，包含3个键值对，并打印出来
	mapValue := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	fmt.Println(mapValue)

	// 练习9：使用for循环从1循环到10，并打印出每个数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习10：使用if语句判断一个整数是否为偶数，如果是则打印“偶数”，否则打印“奇数”
	num2 := 16
	if num2%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}