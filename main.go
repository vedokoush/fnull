package main

import (
	"fmt"
	"fnull/server"
	"fnull/utils"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  fnull send <path/to/file_or_folder>")
	fmt.Println("  fnull receive <fnull-link>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  fnull send secret.txt")
	fmt.Println("  fnull receive I3gan5jQ`")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "send":
		if len(os.Args) < 3 {
			fmt.Println("missing path to file or folder.")
			printUsage()
			return
		}

		sharedPath := os.Args[2]

		info, err := os.Stat(sharedPath)
		if err != nil {
			fmt.Printf("path %s does not exist: %v\n", sharedPath, err)
			return
		}

		isFile := !info.IsDir()
		token := utils.GenerateToken(8)
		localIP := utils.GetLocalIP()
		publicIP := utils.GetPublicIP()

		server.StartServer(sharedPath, isFile, token, 8000, localIP, publicIP)

	case "receive":
		if len(os.Args) < 3 {
			fmt.Println("Err: missing link")
			printUsage()
			return
		}
		link := os.Args[2]
		err := utils.DownloadFromLink(link)
		if err != nil {
			fmt.Printf("Download failed: %v\n", err)
			return
		}
		fmt.Println("Download completed successfully.")

	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printUsage()
	}
}
