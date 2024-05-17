package main

import (
	"fmt"
	"net/http"

	"github.com/QiuMatthew/website-backend/utils/example"
	"github.com/QiuMatthew/website-backend/utils/fileupload"
)

func main() {
	fmt.Println(example.Hello())

	http.HandleFunc("/upload", fileupload.UploadHandler)
	http.ListenAndServe(":8080", nil)
}
