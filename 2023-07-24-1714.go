package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：判断一个整数是奇数还是偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 练习3：将字符串中的字符全部转为大写
	str := "hello, world!"
	fmt.Println(strings.ToUpper(str))

	// 练习4：实现一个简单的HTTP服务器并监听3000端口
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":3000", nil)
}
