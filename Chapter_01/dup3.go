/***********************************
*FileName       :dup3.go
*Author         :LiuYang
*Create Time    :2020/11/11 22:04
*Author's Email :liuyang__work@163.com
************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _,line := range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}
	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}

	}
}