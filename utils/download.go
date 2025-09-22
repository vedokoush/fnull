package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"archive/zip"
)

func DownloadFromLink(link string) error {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("invalid link: %w", err)
	}

	path := strings.Split(parsedURL.Path, "/")
	if len(path) < 2 {
		return fmt.Errorf("invalid format")
	}
	token := path[len(path)-1]
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	isFolder := false
	if strings.Contains(link, "/folder/") {
		isFolder = true
	}

	var downloadURL string
	if isFolder {
		downloadURL = fmt.Sprintf("%s/download/%s.zip", baseURL, token)
	} else {
		downloadURL = link
	}


	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download from %s: %w", downloadURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	if isFolder {
		tmpFile, err := os.CreateTemp("", "fnull-*.zip")
		if err != nil {
			return fmt.Errorf("failed to create temp file: %w", err)
		}
		defer os.Remove(tmpFile.Name())

		_, err = io.Copy(tmpFile, resp.Body)
		if err != nil {
			return fmt.Errorf("failed to copy response body to temp file: %w", err)
		}
		tmpFile.Close()
		err = Unzip(tmpFile.Name(), "./")
		if err != nil {
			return fmt.Errorf("failed to unzip file: %w", err)
		}


	} else {
		filename := "downloaded_file"

		contentDisposition := resp.Header.Get("Content-Disposition")
		if contentDisposition != "" {
			parts := strings.Split(contentDisposition, "filename=")
			if len(parts) > 1 {
				filename = strings.Trim(parts[1], "\"")
			}
		}

		outFile, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w", filename, err)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
		fmt.Printf("Downloaded file: %s\n", filename)
	}

	return nil
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
