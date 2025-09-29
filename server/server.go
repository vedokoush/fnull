package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mdp/qrterminal/v3"
)

var (
	sharedPath string
	isFile     bool
	token      string
)

func StartServer(path string, file bool, tkn string, port int, localIP string, publicIP string) {
	sharedPath, isFile, token = path, file, tkn

	relayRegisterURL = "wss://fnull.shouko.site/register"
	startTunnel()

	link := fmt.Sprintf("https://fnull.shouko.site/%s/", token)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Your Token (session id) is: %s\n\n", token)

	config := qrterminal.Config{
		Level:      qrterminal.M,
		Writer:     os.Stdout,
		QuietZone:  1,
		HalfBlocks: true,
	}
	qrterminal.GenerateWithConfig(link, config)

	fmt.Println()
	fmt.Println("How to receive:")
	fmt.Println("  On the other computer run:")
	fmt.Printf("    fnull receive %s\n\n", token)
	fmt.Println("  Or you can download directly from the link in a browser:")
	fmt.Printf("    %s\n\n", link)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
