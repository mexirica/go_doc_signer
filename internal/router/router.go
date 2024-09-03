package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/go_doc_signer/internal/handlers"
)

func Initialize() {
	r := gin.Default()

	r.POST("/sign", handlers.SignerHandler)
	r.POST("/verify", handlers.VerifyHandler)

	r.Run(":8080")
}
