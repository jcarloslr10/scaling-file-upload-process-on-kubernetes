package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	// _ "net/http/pprof"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const BUFFER_SIZE = 32 << 10     // 32KB
const HEADER_SIZE = 512          // 512B
const MAX_UPLOAD_SIZE = 32 << 20 // 32MB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE+1024)
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	part, err := reader.NextPart()
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if part.FormName() != "file" {
		http.Error(w, "File field is expected", http.StatusBadRequest)
		return
	}

	filename := part.FileName()
	buf := bufio.NewReader(part)
	sniff, _ := buf.Peek(HEADER_SIZE)
	contentType := http.DetectContentType(sniff)
	if contentType != "application/pdf" {
		http.Error(w, "File format is not allowed. Please upload a pdf", http.StatusBadRequest)
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUuid := uuid.New()
	newFileName := strings.Replace(newUuid.String(), "-", "", -1)
	tempFile, err := os.Create(fmt.Sprintf("./uploads/%s%s", newFileName, filepath.Ext(filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer tempFile.Close()

	limitedReader := io.MultiReader(buf, io.LimitReader(part, MAX_UPLOAD_SIZE-(HEADER_SIZE-1)))
	copyBuffer := make([]byte, BUFFER_SIZE)
	written, err := io.CopyBuffer(tempFile, limitedReader, copyBuffer)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if written > MAX_UPLOAD_SIZE {
		os.Remove(tempFile.Name())
		http.Error(w, "Maximum file size exceeded", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Upload successful"))
	return
}

func main() {
	// pprofMux := http.DefaultServeMux

	// go func() {
	// 	log.Println(http.ListenAndServe(":8081", pprofMux))
	// }()

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/", indexHandler)
	apiMux.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(":8080", apiMux))
}
