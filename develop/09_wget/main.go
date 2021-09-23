package main

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

func main() {
	var urel, folderPath string

	urel = os.Args[1]
	folderPath = "./testDirectory/"
	maxDepth := 1

	Download(urel, folderPath, maxDepth)
}

func Download(urel, folderPath string, maxDepth int) {
	if maxDepth < 0 {
		return
	}

	if _, e := url.ParseRequestURI(urel); e != nil {
		return
	}

	fmt.Println("Download: ", urel)

	res, err := http.Get(urel)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	filename := getFilename(urel, res.Header)
	out, err := os.Create(folderPath + filename)
	if err != nil {
		return
	}
	defer out.Close()

	out.Write(data)

	links := getParseLinks(data)
	var wg sync.WaitGroup
	for _, sublink := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			Download(link, folderPath, maxDepth-1)
		}(sublink)
	}
	wg.Wait()
}

func getFilename(url string, header http.Header) string {
	mediaType := createFilename(header)
	baseName := path.Base(url)

	if mediaType == "" {
		return "Error mediaType"
	}

	newName := Cut(Cut(baseName, "#", 0), "?", 0)

	if path.Ext(newName) == "" && mediaType != "" {
		return newName + "." + mediaType
	}

	return newName
}

func createFilename(header http.Header) string {
	contentType := header.Get("Content-Type")
	mType, _, _ := mime.ParseMediaType(contentType)
	mediaType := Cut(mType, "/", 1)
	return mediaType
}

func Cut(s, sep string, i int) string {
	if strings.Contains(s, sep) {
		return strings.Split(s, sep)[i]
	}
	return s
}

func getParseLinks(data []byte) []string {
	reg := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	result := reg.FindAll(data, -1)

	subUrls := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		subUrls = append(subUrls, string(result[i]))
	}

	return subUrls
}
