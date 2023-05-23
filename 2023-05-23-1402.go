package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子，使得每次运行结果不同

	// 练习1：生成一个长度为10的随机整数数组，并打印出来
	var arr1 [10]int
	for i := 0; i < 10; i++ {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println(arr1)

	// 练习2：从上述数组中找出最大值和最小值，并打印出来
	var max, min int = arr1[0], arr1[0]
	for _, val := range arr1 {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Printf("最大值为：%d，最小值为：%d\n", max, min)

	// 练习3：使用冒泡排序算法对上述数组进行从小到大排序，并打印出来
	for i := 0; i < len(arr1)-1; i++ {
		for j := 0; j < len(arr1)-1-i; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	fmt.Println(arr1)

	// 练习4：定义一个Person结构体，包含name、age、gender三个字段，实例化一个Person对象并打印出来
	type Person struct {
		name   string
		age    int
		gender string
	}
	p := Person{name: "Tom", age: 18, gender: "male"}
	fmt.Println(p)
}