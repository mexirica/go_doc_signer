package main

import (
	"github.com/mexirica/go_doc_signer/internal/router"
	signer "github.com/mexirica/go_doc_signer/pkg"
)

func main() {
	signer.InitializeKeys()
	router.Initialize()
}
