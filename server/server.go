package server

import (
	"fmt"
	"net/http"
	"strings"
	"os"
	"github.com/mdp/qrterminal/v3"
)

var (
	sharedPath string
	isFile     bool
	token      string
)

func StartServer(path string, file bool, tkn string, port int, localIP string, publicIP string) {
	sharedPath, isFile, token = path, file, tkn

	link := fmt.Sprintf("http://%s:%d/%s/", localIP, port, token)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Your Token is: %s\n\n", token)

	config := qrterminal.Config{
		Level:     qrterminal.M,  
		Writer:    os.Stdout,
		QuietZone: 1,            
		HalfBlocks: true,       
	}
	qrterminal.GenerateWithConfig(link, config)



	fmt.Println("\nOn the other computer, please run:")
	fmt.Printf("fnull --download %s\n\n", link)

	fmt.Println("Or you can download direct from link: ")
	fmt.Printf("%s\n\n", link)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}