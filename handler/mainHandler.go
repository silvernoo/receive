package handler

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	Path string
)

func Receive(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	pPath := filepath.Join(Path, context.Request.FormValue("path"))
	filename := header.Filename
	if _, err := os.Stat(pPath); os.IsNotExist(err) {
		_ = os.MkdirAll(pPath, os.ModeDir)
	}
	out, err := os.Create(filepath.Join(pPath, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filepath.Join(pPath, filename))
}
