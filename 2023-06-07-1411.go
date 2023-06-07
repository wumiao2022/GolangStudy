package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将字符串转换为整型
	numStr := "123"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d\n", numInt)

	// 将整型转换为字符串
	numInt2 := 456
	numStr2 := strconv.Itoa(numInt2)
	fmt.Printf("%s\n", numStr2)

	// 将字符串转换为浮点型
	floatStr := "3.14"
	floatNum, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%f\n", floatNum)

	// 将浮点型转换为字符串
	floatNum2 := 2.71828
	floatStr2 := strconv.FormatFloat(floatNum2, 'f', 2, 64)
	fmt.Printf("%s\n", floatStr2)

	// 判断字符串是否可以转换为整型
	testStr1 := "hello, world!"
	_, err1 := strconv.Atoi(testStr1)
	if err1 != nil {
		fmt.Println(err1)
	}

	// 判断字符串是否可以转换为浮点型
	testStr2 := "3.14.159"
	_, err2 := strconv.ParseFloat(testStr2, 64)
	if err2 != nil {
		fmt.Println(err2)
	}
}