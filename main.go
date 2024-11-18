package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/films/:name", StreamFilm)

	r.Run(":8000")
}

func StreamFilm(c *gin.Context) {
	filmDir := "./films"

	f := c.Param("name")
	if f == "" {
		c.JSON(400, gin.H{"error": "filname field required in params"})
		return
	}

	fp := filepath.Join(filmDir, f+".mp4")
	fmt.Println(fp)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": "Movie does not exist"})
		return 
	}

	c.Header("Content-Type", "video/mp4")
	c.File(fp)

}