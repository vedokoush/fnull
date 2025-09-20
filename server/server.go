package server

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	sharedPath string
	isFile     bool
	token      string
)

func StartServer(path string, file bool, tkn string, port int, localIP string) {
	sharedPath, isFile, token = path, file, tkn

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("🚀 fnull Phase 1 - Local File Sharing (Go)")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("📁 Sharing: %s\n", sharedPath)
	fmt.Printf("📄 Type: %s\n", map[bool]string{true: "File", false: "Directory"}[isFile])
	fmt.Printf("🔑 Token: %s\n", token)
	fmt.Printf("🌐 Local URL: http://%s:%d/%s/\n", localIP, port, token)
	fmt.Printf("🏠 Localhost: http://localhost:%d/%s/\n", port, token)

	if !isFile {
		fmt.Printf("📦 Direct folder download: http://%s:%d/%s/download.zip\n", localIP, port, token)
	} else {
		fmt.Printf("⬇️  Direct file download: http://%s:%d/%s/download\n", localIP, port, token)
	}
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("❌ Error starting server: %v\n", err)
	}
}