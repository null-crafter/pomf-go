package main

import (
	"log"
	"net/http"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 128 // 128MB
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 128MB in size", http.StatusBadRequest)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload.php", uploadHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
