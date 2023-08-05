package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前的日期和时间
	fmt.Println(time.Now())

	// 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 声明一个切片并打印其中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 使用if-else语句判断一个数字是偶数还是奇数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 声明一个map并打印其中的键值对
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	for key, value := range m {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
```

记得多练习哦！