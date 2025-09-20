package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func serveSingleFile(w http.ResponseWriter, r *http.Request, download bool) {
	file, err := os.Open(sharedPath)
	if err != nil {
		http.Error(w, "Error opening file", 500)
		return
	}
	defer file.Close()

	stat, _ := file.Stat()
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", stat.Name()))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error sending file", 500)
	}
}