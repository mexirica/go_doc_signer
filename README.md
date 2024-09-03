# Go Document Signer

Go Document Signer is a Go-based project designed for digital signing and verification of documents using RSA cryptography. It provides functionalities to sign documents, and verify their signatures.

It comes with a little web-api made using Gin to see how it works.


## Features

- **Document Signing**: Sign documents using a private RSA key.
- **Signature Verification**: Verify the authenticity of signed documents using a public RSA key.

## Installation

To get started with this project, follow these steps:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/mexirica/go_doc_signer.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd go_doc_signer
    ```

3. **Install the required dependencies:**

    ```bash
    go mod tidy
    ```

4. **Run the application:**
    ```bash
    go run main.go
    ```


## HTTP API

The project provides HTTP handlers for signing and verifying documents.

- **POST /sign**

Form Parameters:

- file: The document to be signed.

Response: base64-encoded-signature

- **POST /verify**

Form Parameters:

- file: The document to be signed.
- signature: base64-encoded-signature
  
Response: bool
