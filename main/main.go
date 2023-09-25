package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const TargetPath = "/home/UploadDir"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving file", err)

		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %v\n", handler.Filename)
	fmt.Printf("File Size: %v\n", handler.Size)
	fmt.Printf("MIME Header %v\n", handler.Header)
	tempFile, err := os.CreateTemp(TargetPath, "*-"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Upload File\n")
}

func SetupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Serving Upload File")
	SetupRoutes()
}
