package utils

import (
	"fmt"
	"io"
	"mime/multipart"
)

// ConvertFileToBytes Converts a multipart FileHeader to bytes.
//
// It takes a pointer to a multipart.FileHeader as input and returns a slice of bytes
// representing the entire contents of the file, along with an error if any occurs during
// the conversion process.
//
// The function reads the entire file into memory, so it may be inefficient for very large files.
//
// Example usage:
//
//	fileHeader, err := formFile("fileField")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	data, err := convertFileToBytes(fileHeader)
//	if err != nil {
//		log.Fatal(err)
//	}
func ConvertFileToBytes(file *multipart.FileHeader) ([]byte, error) {
	buffer := make([]byte, file.Size)

	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("Error to open the file: %w", err)
	}
	defer f.Close()

	_, err = io.ReadFull(f, buffer)
	if err != nil {
		return nil, fmt.Errorf("Error to read the file: %w", err)
	}

	return buffer, nil
}

// ReadFileInParts Reads a multipart file in parts using streaming and buffering.
//
// It takes a pointer to a multipart.FileHeader as input and returns a slice of byte slices,
// each representing a part of the file contents, along with an error if any occurs during
// the conversion process.
//
// The function uses buffered reader to efficiently read the file in chunks.
//
// Example usage:
//
//	fileHeader, err := formFile("fileField")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	parts, err := ReadFileInParts(fileHeader)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, part := range parts {
//		fmt.Println(len(part))
//	}
func ReadFileInParts(file *multipart.FileHeader) ([][]byte, error) {
	bufferSize := 1024 * 1024 // 1MB buffer size
	buffer := make([]byte, bufferSize)

	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer f.Close()

	var parts [][]byte
	for {
		n, err := f.Read(buffer)
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			break
		}
		parts = append(parts, buffer[:n])
	}

	if err != io.EOF {
		return nil, fmt.Errorf("Unexpected error reading file: %w", err)
	}

	return parts, nil
}
