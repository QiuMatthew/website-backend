package main

import (
	"fmt"
	"net/http"

	"github.com/QiuMatthew/website-backend/discretelog"
	"github.com/QiuMatthew/website-backend/utils/example"
	"github.com/QiuMatthew/website-backend/utils/fileupload"
)

func main() {
	fmt.Println(example.Hello())

	http.HandleFunc("/upload", fileupload.UploadHandler)
	http.HandleFunc("/discrete-log", discretelog.DiscreteLogHandler)
	http.ListenAndServe(":8080", nil)
}
