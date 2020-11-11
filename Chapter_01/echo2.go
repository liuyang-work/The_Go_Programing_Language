/***********************************
*FileName       :echo2.go
*Author         :LiuYang
*Create Time    :2020/11/11 19:59
*Author's Email :liuyang__work@163.com
************************************/
//echo1 输出其命令行参数
package main

import "fmt"
import "os"

func main() {
	s, sep := "", ""
	for _,arg:= range os.Args[1:]{
		s += sep + arg
		sep = " "
		fmt.Println(s)
	}
}
