package templater

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trimPath(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx != -1 {
		path = path[:idx]
	}
	return path
}

func templateMarkdowntoHTML(FileLocations string) {
	mdFile, ok := os.Open(FileLocations)
	if ok != nil {
		panic(ok)
	}
	scanner := bufio.NewScanner(mdFile)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	mdFile.Close()
	if _, err := os.Stat("./html"); errors.Is(err, os.ErrNotExist) {
		panic("html directory is not exists")
	}
	FileLocations = strings.Replace(FileLocations, "markdown", "./html", 1)
	fmt.Println(FileLocations)
	htmlFileName := FileLocations[:len(FileLocations)-2] + "html"
	// os.MkdirAll(folderPath, os.ModePerm)
	// fmt.Println(trimPath(htmlFileName))
	os.MkdirAll(trimPath(htmlFileName), 0700)
	// createEmptyFile := func(name string) {
	// 	d := []byte("")
	// 	check(os.WriteFile(name, d, 0644))
	// }
	// createEmptyFile(htmlFileName)

	// if _, err := os.Stat(htmlFileName); errors.Is(err, os.ErrNotExist) {
	// panic("html directory is not exists")
	// createEmptyFile(htmlFileName)
	// }
	htmlFile, ok := os.Create(htmlFileName)
	if ok != nil {
		panic(ok)
	}
	for _, lineIterator := range text {
		md := []byte(lineIterator)
		rendered := markdown.ToHTML(md, nil, nil)
		_, ok := htmlFile.WriteString(string(rendered))
		if ok != nil {
			panic(ok)
		}
	}
	htmlFile.Close()
}

func find(content, ext string) []string {
	var a []string
	filepath.WalkDir(content, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func Run() {
	for _, s := range find("./markdown", ".md") {
		templateMarkdowntoHTML(s)
	}

	// r := gin.Default()
	// // r.LoadHTMLGlob("./templates/html/*.html")
	// // r.LoadHTMLGlob("/home/safa/works/go/goblog/templates/html/post/wine-data-mining/Turkish/*.html")
	// r.LoadHTMLGlob("/home/safa/works/go/goblog/html/**/*")
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	// r.GET("/aboutme", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "about.html", nil)
	// })
	// r.GET("/findprune", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "post/find-prune.html", nil)
	// })
	// r.GET("/postgresql-localhost", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "post/Postgresql-Localhost.html", nil)
	// })
	// r.GET("/replacement-algorithm", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "post/Replacement-Algorithm.html", nil)
	// })
	// r.GET("/virtualhost-postfix", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "post/virtualhostpostfix.html", nil)
	// })
	// r.Run(":3000")

	// err := tmpl.ExecuteTemplate(w, "page.html", p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
