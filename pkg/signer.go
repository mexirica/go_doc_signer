package pkg

import (
    "crypto/sha256"
    "encoding/hex"
    "crypto/rand"
    "crypto/rsa"
    "crypto"
    "encoding/base64"
)

// Calculates the SHA256 hash of a document.
//
// It takes a byte slice representing the document and returns the
// hexadecimal string representation of the hash.
//
// Example usage:
//
//    doc := []byte("Hello, Go!")
//    hash := hashDocument(doc)
//    fmt.Println(hash)
func HashDocument(doc []byte) string {
    hash := sha256.New()
    hash.Write(doc)
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}

// Signs a document hash using an RSA private key.
//
// It takes the document hash as a string and the private key as a
// pointer to an rsa.PrivateKey. It returns the base64-encoded string
// representation of the signature and an error, if any.
//
// Example usage:
//
//    docHash := "exampleHash"
//    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
//    signature, err := signDocument(docHash, privateKey)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(signature)
func SignDocument(docHash string, privateKey *rsa.PrivateKey) (string, error) {
    hash := sha256.Sum256([]byte(docHash))
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(signature), nil
}

// Verifies the signature of a document hash using an
// RSA public key.
//
// It takes the document hash as a string, the signature as a base64-encoded
// string, and the public key as a pointer to an rsa.PublicKey. It returns a
// boolean indicating whether the signature is valid and an error, if any.
//
// Example usage:
//
//    docHash := "exampleHash"
//    signature := "exampleSignature"
//    publicKey, _ := rsa.GenerateKey(rand.Reader, 2048).Public().(*rsa.PublicKey)
//    valid, err := verifySignature(docHash, signature, publicKey)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(valid)
func VerifySignature(docHash string, signatureStr string, publicKey *rsa.PublicKey) (bool, error) {
    hash := sha256.Sum256([]byte(docHash))
    signature, err := base64.StdEncoding.DecodeString(signatureStr)
    if err != nil {
        return false, err
    }
    err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
    return err == nil, nil
}