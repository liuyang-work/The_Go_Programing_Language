/***********************************
*FileName       :dup2.go
*Author         :LiuYang
*Create Time    :2020/11/11 21:42
*Author's Email :liuyang__work@163.com
************************************/
//dup2打印输入中多次出现的行的个数和文本
//他从stdin或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {

		coutlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			coutlines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {

			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}

func coutlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//注意：忽略input.Err()中可能的错误
}
