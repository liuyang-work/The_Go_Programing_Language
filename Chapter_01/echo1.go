/***********************************
*FileName       :echo1.go
*Author         :LiuYang
*Create Time    :2020/11/11 19:47
*Author's Email :liuyang__work@163.com
************************************/
//echo1 输出其命令行参数
package main

import "fmt"
import "os"

func main() {
	var s, sep string
	for i := 1; i < len(os.Args) ;i ++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
