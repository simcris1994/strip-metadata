package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"strip-metadata/processing"
)

func main() {
	router := gin.Default()

	router.POST("/upload", upload)

	router.Run()
}

func upload(c *gin.Context) {
	// fetch file
	file, _ := c.FormFile("file")

	// save file
	dest := "./images/" + file.Filename
	strippedDest := "./images/stripped_" + file.Filename
	c.SaveUploadedFile(file, dest)

	// call processing module to strip the image of exif data
	err := processing.StripMetadata(dest, strippedDest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "File processing failed"})
	}

	// delete original and processed file
	os.Remove(dest)
	defer os.Remove(strippedDest)

	// return file
	c.File(strippedDest)
}
