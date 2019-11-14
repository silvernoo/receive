package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	Path string
)

func Receive(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	filename := header.Filename
	print(filename)
	out, err := os.Create(filepath.Join(Path, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
