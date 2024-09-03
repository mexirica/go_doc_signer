package handlers

import (
	"github.com/mexirica/go_doc_signer/internal/models"
	"github.com/mexirica/go_doc_signer/pkg/signer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignerHandler(c *gin.Context) {
	file, h, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	signature, err := signer.SignDocument(h, signer.PrivateKey)

	response := models.Response{Response: signature}

	c.JSON(200, response)
}

func VerifyHandler(c *gin.Context) {
	signature := c.Request.FormValue("signature")
	if signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "signature required"})
	}
	file, h, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	isValid := signer.VerifySignature(h, signature, signer.PublicKey)

	var response models.Response
	if isValid {
		response = models.Response{
			Response: "The document is the same",
		}
	} else {
		response = models.Response{
			Response: "The document has been modified",
		}
	}

	c.JSON(200, response)
}
