package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"receive/handler"
	"receive/initRouter"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	f string
)

//example.txt
func init() {
	flag.StringVar(&f, "f", "", "Set configuration file")
}

func main() {
	flag.Parse()
	if f != "" {
		download()
	} else {
		if os.Getenv("DEBUG") == "debug" {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		handler.Path, _ = os.Getwd()
		router := initRouter.SetupRouter()
		_ = router.Run(":80")
	}
}

func download() {
	f, e := os.Open(f)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		url := strings.Split(s.Text(), " ")[0]
		resp, _ := http.Get(url)
		sprint := strings.Split(url, "/")
		putFile(resp.Body, sprint[len(sprint)-1], strings.Split(s.Text(), " ")[1])
		resp.Body.Close()
	}
}

func putFile(read io.ReadCloser, filename string, remoteLocaltion string) {
	println(remoteLocaltion)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(filename))
	err := writer.WriteField("path", remoteLocaltion)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, read)
	writer.Close()
	r, _ := http.NewRequest("PUT", "http://127.0.0.1/receive", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)
}
