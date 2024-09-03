package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/mexirica/go_doc_signer/internal/utils"
	"mime/multipart"
)

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func InitializeKeys() {
	PrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	PublicKey = &PrivateKey.PublicKey
}

// SignDocument Signs a document hash using an RSA private key.
//
// It takes the document hash as a string and the private key as a
// pointer to an rsa.PrivateKey. It returns the base64-encoded string
// representation of the signature and an error, if any.
//
// Example usage:
//
//	docHash := "exampleHash"
//	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
//	signature, err := signDocument(docHash, privateKey)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(signature)
func SignDocument(doc *multipart.FileHeader, privateKey *rsa.PrivateKey) (string, error) {
	docBytes, err := utils.ConvertFileToBytes(doc)
	hash := sha256.Sum256(docBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySignature Verifies the signature of a document hash using an
// RSA public key.
//
// It takes the document hash as a string, the signature as a base64-encoded
// string, and the public key as a pointer to an rsa.PublicKey. It returns a
// boolean indicating whether the signature is valid and an error, if any.
//
// Example usage:
//
//	docHash := "exampleHash"
//	signature := "exampleSignature"
//	publicKey, _ := rsa.GenerateKey(rand.Reader, 2048).Public().(*rsa.PublicKey)
//	valid, err := verifySignature(docHash, signature, publicKey)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(valid)
func VerifySignature(file *multipart.FileHeader, signatureStr string, publicKey *rsa.PublicKey) bool {
	bytesdoc, err := utils.ConvertFileToBytes(file)
	hash := sha256.Sum256(bytesdoc)
	signature, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return false
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	return err == nil
}
