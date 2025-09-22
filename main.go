package main

import (
	"flag"
	"fmt"
	"os"
	"fnull/server"
	"fnull/utils"
)

func main() {
	port := flag.Int("port", 8000, "Port to use (default 8000)")
	download := flag.String("download", "", "Download from fnull link")
	flag.Parse()

	if *download != "" {
		utils.DownloadFromLink(*download)
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		fmt.Println("Usage: go run main.go /path/to/file_or_folder [--port 9000]")
		return
	}

	sharedPath := flag.Arg(0)
	info, err := os.Stat(sharedPath)
	if err != nil {
		fmt.Printf("Error: path %s does not exist\n", sharedPath)
		return
	}

	isFile := !info.IsDir()
	token := utils.GenerateToken(8)
	localIP := utils.GetLocalIP()
	publicIP := utils.GetPublicIP()

	server.StartServer(sharedPath, isFile, token, *port, localIP, publicIP)
}
