package server

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func serveFolderAsZip(w http.ResponseWriter, r *http.Request) {
	zipName := filepath.Base(sharedPath) + ".zip"

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	filepath.Walk(sharedPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(sharedPath, path)
		f, err := zipWriter.Create(rel)
		if err != nil {
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(f, srcFile)
		return err
	})
}