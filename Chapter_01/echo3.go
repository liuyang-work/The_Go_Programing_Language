/***********************************
*FileName       :echo3.go
*Author         :LiuYang
*Create Time    :2020/11/11 20:11
*Author's Email :liuyang__work@163.com
************************************/
//echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
	//fmt.Println(os.Args[1:])
}