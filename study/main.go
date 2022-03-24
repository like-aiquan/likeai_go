// Package main by chenaiquan<like.aiquan@qq.com> create on 2021/9/29 20.46
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	isEject    bool
	confPath   string
	scriptName string
)

func main() {
	// 编译后生成的文件 可读取参数  eject   database-script
	fmt.Println(os.Executable())
	fmt.Println(RelativePath("conf.ini"))
	flag.StringVar(&confPath, "c", RelativePath("conf.ini"), "配置文件路径")
	flag.BoolVar(&isEject, "eject", false, "导出内置静态资源")
	flag.StringVar(&scriptName, "database-script", "", "运行内置数据库助手脚本")
}

// RelativePath 获取相对可执行文件的路径
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}
