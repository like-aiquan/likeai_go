// Package main by chenaiquan<like.aiquan@qq.com> create on 2021/9/30 17.35
package main

import (
	"fmt"
	"github.com/likeai/study/http/CG"
	"net/http"
)

func main() {
	r := CG.New()
	r.GET("/cg/hello", helloHandler)
	r.GET("/cg/like", queryName)
	r.POST("/cg/post/test", postHandler)
	_ = r.Run(":8080")
}

func postHandler(c *CG.Context) {
	fmt.Println(c.PostForm("name"))
	c.Json(http.StatusOK, CG.H{
		"username": "qqai",
		"age":      "12",
	})
}

func helloHandler(c *CG.Context) {
	fmt.Println(c.Query("name"))
	_, _ = fmt.Fprintf(c.Response, c.Path+"hello world!")
}

func queryName(c *CG.Context) {
	name := c.Query("name")
	_, _ = fmt.Fprintf(c.Response, "hello %s", name)
}
