package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println("当前时间：", time.Now())

	// 循环打印数字1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 定义一个字符串变量，并逐个打印出来
	s := "Hello, world!"
	for _, c := range s {
		fmt.Println(string(c))
	}

	// 定义一个字典变量，输出键值对
	dict := make(map[string]string)
	dict["apple"] = "苹果"
	dict["banana"] = "香蕉"
	dict["orange"] = "橙子"
	for k, v := range dict {
		fmt.Println(k, "的中文名是", v)
	}
}