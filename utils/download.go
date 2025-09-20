package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFromLink(downloadURL string) {
	fmt.Printf("🔗 Attempting to access: %s\n", downloadURL)
	resp, err := http.Get(downloadURL)
	if err != nil {
		fmt.Printf("❌ Error accessing link: %v\n", err)
		return
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if strings.Contains(ct, "text/html") {
		fmt.Println("📂 Directory listing found or HTML content!")
		fmt.Printf("💡 Browse: %s\n", downloadURL)
		fmt.Printf("📦 Try ZIP: %sdownload.zip\n", downloadURL)
	} else {
		u, _ := url.Parse(downloadURL)
		name := filepath.Base(u.Path)
		if name == "" {
			name = "downloaded_file"
		}
		file, err := os.Create(name)
		if err != nil {
			fmt.Printf("❌ Error creating file: %v\n", err)
			return
		}
		defer file.Close()
		n, _ := io.Copy(file, resp.Body)
		fmt.Printf("✅ Downloaded: %s (%d bytes)\n", name, n)
	}
}