package fileupload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// solve cors error
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	err := r.ParseMultipartForm(10 << 20) // up to 10 MegaBytes
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// create folder for uploaded files
	savePath := "../assets/easy-slide/uploads"
	err = os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// create new empty file
	//dst, err := os.Create(filepath.Join(savePath, handler.Filename))
	dst, err := os.Create(filepath.Join(savePath, "slide.md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer dst.Close()

	// copy the uploaded content to the newly created file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "file uploaded successfully: %s\n", handler.Filename)
}
