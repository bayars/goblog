package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

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
	FileLocations = strings.Replace(FileLocations, "markdown", "./templates/html", 1)
	fmt.Println(FileLocations)
	htmlFileName := FileLocations[:len(FileLocations)-2] + "html"
	// os.MkdirAll(folderPath, os.ModePerm)
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

func main() {
	for _, s := range find("./markdown", ".md") {
		templateMarkdowntoHTML(s)
	}

	r := gin.Default()
	// r.LoadHTMLGlob("./templates/html/*.html")
	r.LoadHTMLGlob("/home/safa/works/go/goblog/templates/html/post/wine-data-mining/Turkish/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "README.html", nil)
	})
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	r.Run(":3000")

	// err := tmpl.ExecuteTemplate(w, "page.html", p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
