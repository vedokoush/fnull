package server

import (
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) == 0 || parts[0] != token {
		http.Error(w, "Invalid token or path not found", 404)
		return
	}

	if isFile {
		if len(parts) == 1 || (len(parts) == 2 && parts[1] == "download") {
			serveSingleFile(w, r, true)
			return
		}
	} else {
		if len(parts) == 2 && parts[1] == "download.zip" {
			serveFolderAsZip(w, r)
			return
		}
		http.StripPrefix("/"+token+"/", http.FileServer(http.Dir(sharedPath))).ServeHTTP(w, r)
	}
}