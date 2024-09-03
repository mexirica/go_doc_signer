package main

import (
	"github.com/mexirica/go_doc_signer/internal/router"
	"github.com/mexirica/go_doc_signer/pkg/signer"
)

func main() {
	signer.InitializeKeys()
	router.Initialize()
}
