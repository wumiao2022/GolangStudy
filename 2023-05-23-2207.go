package main

import "fmt"

func main() {
	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")
	
	// 2. 打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 3. 判断一个数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "不是偶数")
	}
	
	// 4. 计算两个数的和
	a, b := 5, 3
	fmt.Println("a + b =", a+b)
	
	// 5. 判断一个字符串是否包含另一个字符串
	str1 := "hello world"
	str2 := "lo w"
	if contain := strings.Contains(str1, str2); contain {
		fmt.Println(str1, "包含", str2)
	} else {
		fmt.Println(str1, "不包含", str2)
	}
	
	// 6. 定义一个结构体并初始化 
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Tom", Age: 20}
	fmt.Println(p)
	
	// 7. 使用 map 存储并访问键值对
	scoreMap := make(map[string]int)
	scoreMap["Tom"] = 90
	scoreMap["John"] = 80
	fmt.Println(scoreMap["Tom"])
	
	// 8. 定义一个切片并进行切片操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[1:3])
	
	// 9. 使用 defer 关键字延迟函数的执行
	defer func() {
		fmt.Println("deferred function")
	}()
	
	// 10. 自定义一个错误类型并抛出错误
	err := errors.New("custom error")
	if err != nil {
		panic(err)
	}
}