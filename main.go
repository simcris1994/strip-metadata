package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", upload)

	router.Run("localhost:8080")
}

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dest := "./images/" + file.Filename
	c.SaveUploadedFile(file, dest)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
