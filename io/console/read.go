// Package console by chenaiquan<like.aiquan@qq.com> create on 2021/10/21 13.30
package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadFromConsole 命令行读
func ReadFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Println("read exit end;")
	for {
		line, _ := reader.ReadString('\n')
		// 读入有换行符
		line = strings.ReplaceAll(line, "\n", "")
		if strings.EqualFold("exit", line) {
			break
		}
		s += line
	}
	return s
}
