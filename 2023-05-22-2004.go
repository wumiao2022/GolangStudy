package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 将华氏温度转换为摄氏温度
	var fahrenheit float64 = 100.0
	var celsius float64 = (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f华氏度 = %.2f摄氏度\n", fahrenheit, celsius)

	// 3. 计算圆的面积和周长
	var radius float64 = 5
	var circumference float64 = 2 * 3.14 * radius
	var area float64 = 3.14 * radius * radius
	fmt.Printf("半径为%.2f的圆的周长为%.2f，面积为%.2f\n", radius, circumference, area)

	// 4. 判断一个数是否为素数
	var num int = 17
	var flag bool = true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 5. 实现冒泡排序
	var arr = []int{3, 5, 2, 1, 4}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}