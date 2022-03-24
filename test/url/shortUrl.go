// Package url by chenaiquan<like.aiquan@qq.com> create on 2021/11/11 23.01
package url

import (
	"fmt"
	"google.golang.org/api/urlshortener/v1"
	"html/template"
	"net/http"
)

var templateRoot = template.Must(template.New("rootHtml").Parse(`
<html>
<body>
    <h1>Short_Url</h1>
    {{if .}}{{.}} <br/><br/> {{end}}
    <form action="/short" class="post">
        <input type="text" name="shortUrl">
        <input type="submit" value="short">
    </form>
    <form action="/long" class="post">
        <input type="text" name="long">
        <input type="submit" value="long">
    </form>
</body>
</html>
`))

func Listener() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)
	_ = http.ListenAndServe("localhost:8080", nil)
}

func long(w http.ResponseWriter, r *http.Request) {
	value := "http://test.gl/" + r.FormValue("long")
	service, _ := urlshortener.New(http.DefaultClient)
	longUrl, _ := service.Url.Get(value).Do()
	_ = templateRoot.Execute(w, fmt.Sprintf("long url %s", longUrl.Id))
}

func short(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("short")
	service, _ := urlshortener.New(http.DefaultClient)
	url, _ := service.Url.Insert(&urlshortener.Url{LongUrl: value}).Do()
	_ = templateRoot.Execute(w, fmt.Sprintf("short url %s", url.Id))
}

func root(w http.ResponseWriter, r *http.Request) {
	_ = templateRoot.Execute(w, nil)
}
