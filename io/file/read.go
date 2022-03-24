// Package file by chenaiquan<like.aiquan@qq.com> create on 2021/10/21 13.49
package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ReadFromFile(path string) (s string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// after hock
	defer func() {
		_ = file.Close()
		fmt.Println("文件关闭")
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		s += line
		if err == io.EOF {
			// 最后一行会同时返回 EOF 标志
			return
		}
		if err != nil {
			fmt.Println("error read")
			return
		}
	}
}

// AddAnnotationOnProperty 给 java 文件的类字段添加注解
func AddAnnotationOnProperty(directory string) {
	// 文件加
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
	}
	// 遍历文件
	for _, file := range files {
		// 是否以 .java 结尾
		if strings.HasSuffix(file.Name(), ".java") {
			// 添加注解
			go addAnnotationOnProperty(file, directory)
		}
	}
}

func addAnnotationOnProperty(fileInfo fs.FileInfo, directory string) {
	path := directory + "/" + fileInfo.Name()
	// 读权限操作文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	var s string
	for {
		line, err := reader.ReadString('\n')
		organLine := strings.ReplaceAll(line, "\n", "")
		matched, _ := regexp.MatchString("(\\s+)(private|public)\\s+[A-Z|a-z]+\\s+[A-Za-z]+;", organLine)
		if matched {
			split := strings.Split(organLine, " ")
			filedName := strings.ReplaceAll(split[len(split)-1], ";", "")
			filedName = CaseToSnickedName(filedName)
			line = "@JsonProperty(\"" + filedName + "\")\n" + line
		}

		// append file string
		s += line

		if err == io.EOF {
			// 最后一行会同时返回 EOF 标志
			break
		}
		if err != nil {
			fmt.Println("error read")
			return
		}
	}

	_ = file.Close()

	// 覆盖写模式重新打开文件
	file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)

	// write file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(s)
	err = writer.Flush()
	if err != nil {
		fmt.Printf("write file error : %s\n", err)
		return
	}

	// after hock
	defer func() {
		_ = file.Close()
		fmt.Println("文件关闭")
	}()
}

func CaseToSnickedName(name string) (snickName string) {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return strings.ReplaceAll(buffer.String(), "\r", "")
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		_, _ = b.Write(val)
	case rune:
		_, _ = b.WriteRune(val)
	}
	return b
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); !IsNil(err) {
			fmt.Println(err)
		}
	}()
	_, _ = b.WriteString(s)
	return b
}
