// Package api by chenaiquan<like.aiquan@qq.com> create on 2021/9/30 16.04
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var r = gin.Default()

// CreateRoute 初始化所有路由
func CreateRoute() {
	route()
	groupRoute()
	upload()
	// 开启服务 忽略错误
	_ = r.Run()
}

// 普通路由
func route() {
	r.GET("/", helloHandler)
	r.GET("/user/:name", queryName)
}

func upload() {
	//<form enctype="multipart/form-data" action="http://localhost:8080/upload/only/one" method="post">
	//    <input type="file" name="upload_file" />
	//    <input type="hidden" name="token" value="like_ai" />
	//    <input type="submit" value="upload" />
	//</form>
	r.POST("/upload/only/one", indexHandler)
	r.POST("/upload/multipart/file", multipartUpload)
}

// 路由分组
func groupRoute() {
	// group: v1
	v1 := r.Group("/api/v1")
	v1.GET("/", defaultHandler)
	v1.GET("/other", defaultHandler)
	// group: v2
	v2 := r.Group("/api/v2")
	v2.GET("/", defaultHandler)
	v2.GET("/other", defaultHandler)
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func queryName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func indexHandler(c *gin.Context) {
	file, err := c.FormFile("upload_file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.PostForm("token"))
	fmt.Println(file.Filename)
	// c.saveFile(file)
	c.String(http.StatusOK, "%s upload!", file.Filename)
}

func multipartUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload_file"]
	for _, file := range files {
		log.Println(file.Filename)
	}
	c.String(http.StatusOK, "%d files uploaded!", len(files))
}

func defaultHandler(c *gin.Context) {
	// 解析 json 格式返回
	c.JSON(http.StatusOK, gin.H{
		"path": c.FullPath(),
	})
}
