package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", upload)

	router.Run("localhost:8080")
}

func upload(c *gin.Context) {
	// fetch file
	file, _ := c.FormFile("file")

	// save file
	dest := "./images/" + file.Filename
	c.SaveUploadedFile(file, dest)

	// return file
	c.File(dest)
}
