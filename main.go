package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	filename string // 文件名
	version bool // 版本
	help bool // 帮助
)
func init() {
	flag.StringVar(&filename, "f", "", "请输入文件路径")

	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stdout, `md2html version: 0.1
Usage: md2html [-hvi] [-f filename]
Options:
`)
	flag.PrintDefaults()
}
func main(){
	flag.Parse()

	if filename == "" {
		flag.Usage()
		return
	}

	if help { // 显示帮助信息
		flag.Usage()
		return
	}
	if version { // 显示版本信息
		fmt.Fprintf(os.Stdout, `md2html version: 0.1`)
		return
	}
}