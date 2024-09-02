package internal

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mexirica/go_doc_signer/internal/db"
)

func Initialize() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	parentDir := filepath.Dir(filepath.Dir(wd))

	db.Connect([]interface{}{}, parentDir + "/api.db")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}